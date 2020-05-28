package registry

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Requester represents an HTTP(S) endpoint
type Requester struct {
	hostname      string
	port          int
	insecure      bool
	scheme        string
	url           string
	username      string
	password      string
	httpClient    *http.Client
	authenticator Authenticator
}

// NewRequester creates a new Requester
func NewRequester(hostname string, port int, insecure bool, username string, password string) (r *Requester, err error) {
	err = nil

	scheme := "https"
	if insecure {
		scheme = "http"
	}
	r = &Requester{
		hostname:      hostname,
		port:          port,
		insecure:      insecure,
		scheme:        scheme,
		url:           scheme + "://" + hostname + ":" + strconv.Itoa(port),
		username:      username,
		password:      password,
		httpClient:    &http.Client{},
		authenticator: &NullAuthenticator{},
	}

	request, err := http.NewRequest("GET", r.url+"/v2/", strings.NewReader(""))
	if err != nil {
		err = errors.New("Failed to create request")
		return
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		err = errors.New("Failed to get response")
		return
	}
	if response.Close {
		defer response.Body.Close()
	}
	if response.StatusCode != http.StatusOK {
		challenge := response.Header.Get("Www-Authenticate")

		parts := strings.Split(challenge, " ")
		authType := parts[0]

		var realm *url.URL
		var service string
		data := strings.Split(parts[1], ",")
		for i := 0; i < len(data); i++ {
			keyvalue := strings.Split(data[i], "=")

			if keyvalue[0] == "realm" {
				realm, err = url.Parse(strings.Trim(keyvalue[1], "\""))
				if err != nil {
					return
				}
			}
			if keyvalue[0] == "service" {
				service = strings.Trim(keyvalue[1], "\"")
			}
		}

		if authType == "Basic" {
			r.authenticator = &BasicAuthenticator{
				username: r.username,
				password: r.password,
			}
		} else if authType == "Bearer" {
			r.authenticator, err = NewChallengeAuthenticator(r.hostname, r.port, r.scheme, r.username, r.password, *realm, service)
			if err != nil {
				err = errors.New("Failed to instantiate challenge authenticator")
				return
			}
		}
	}

	return
}
