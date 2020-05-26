package registry

import (
	"net/http"
	"net/url"
)

type Authenticator interface {
	AddAuthentication(request *http.Request) error
}

type NullAuthenticator struct{}

func (a NullAuthenticator) AddAuthentication(request *http.Request) error {
	return nil
}

type BasicAuthenticator struct{
	username string
	password string
}

func (a BasicAuthenticator) AddAuthentication(request *http.Request) error {
	return nil
}

type ChallengeAuthenticator struct{
	hostname      string
	port          int
	scheme        string
	username      string
	password      string
	realm         url.URL
	service       string
	access        string
	token         string
}

func (a ChallengeAuthenticator) AddAuthentication(request *http.Request) error {
	return nil
}

func NewChallengeAuthenticator(hostname string, port int, scheme string, username string, password string, realm url.URL, service string) (a ChallengeAuthenticator, err error) {
	a = ChallengeAuthenticator{
		hostname:      hostname,
		port:          port,
		scheme:        scheme,
		username:      username,
		password:      password,
		realm:         realm,
		service:       service,
	}
	//TODO: get token
	return
}