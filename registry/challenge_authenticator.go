package registry

import (
	"errors"
	"net/url"
)

type ChallengeAuthenticator struct {
	hostname string
	port     int
	insecure bool
	username string
	password string
	realm    *url.URL
	service  string
}

func NewChallengeAuthenticator(hostname string, port int, insecure bool, username string, password string, realm *url.URL, service string) *ChallengeAuthenticator {
	return &ChallengeAuthenticator{
		realm: realm,
		service: service,
	}
}

func (a *ChallengeAuthenticator) Test() error {
	return errors.New("blarg")
}