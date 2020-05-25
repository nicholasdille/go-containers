package registry

import (
	"errors"
)

type BasicAuthenticator struct {
	hostname string
	port     int
	insecure bool
	username string
	password string
}

func NewBasicAuthenticator(hostname string, port int, insecure bool, username string, password string) *BasicAuthenticator {
	return &BasicAuthenticator{
		hostname: hostname,
		port: port,
		insecure: insecure,
		username: username,
		password: password,
	}
}

func (a *BasicAuthenticator) Test() error {
	return errors.New("blarg")
}