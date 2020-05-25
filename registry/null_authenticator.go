package registry

import (
	"errors"
)

type NullAuthenticator struct {
	hostname string
	port     int
	insecure bool
}

func NewNullAuthenticator(hostname string, port int, insecure bool) *NullAuthenticator {
	return &NullAuthenticator{
		hostname: hostname,
		port: port,
		insecure: insecure,
	}
}

func (a *NullAuthenticator) Test() error {
	return errors.New("blarg")
}