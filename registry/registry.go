package registry

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"net/http"
	"net/url"
)

type Challenge struct {
	realm   *url.URL
	service string
}

type Token struct {
	value string
}

type Registry struct {
	hostname string
	port int
	insecure bool
	url string
	username string
	password string
	authType string
	challenge *Challenge
	token *Token
}

func NewRegistry(hostname string, port int, insecure bool, username string, password string) Registry {
	var url = hostname + ":" + strconv.Itoa(port) + "/v2/"
	if insecure {
		url = "http://" + url
	} else {
		url = "https://" + url
	}

	return Registry{
		hostname: hostname,
		port: port,
		insecure: insecure,
		url: url,
		username: username,
		password: password,
	}
}

func (r* Registry) String() string {
	return fmt.Sprintf("hostname=<%s>, port=<%d>, insecure=<%t>, url=<%s>, authType=<%s>, realm=<%s>, service=<%s>", r.hostname, r.port, r.insecure, r.url, r.authType, r.challenge.realm, r.challenge.service)
}

func (r *Registry) Check() (err error) {
	httpClient := &http.Client{}

	request, err := http.NewRequest("GET", r.url, strings.NewReader(""))
	if err != nil {
		return errors.New("Failed to create request")
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return errors.New("Failed to get response")
	}
	if response.Close {
		defer response.Body.Close()
	}
	if response.StatusCode != http.StatusOK {
		challenge := response.Header.Get("Www-Authenticate")

		parts := strings.Split(challenge, " ")
		r.authType = parts[0]

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

		r.challenge = &Challenge{
			realm: realm,
			service: service,
		}

		return
	}

	return nil
}