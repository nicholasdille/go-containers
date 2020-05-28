package registry

import (
	"net/http"
	"net/url"
)

// Authenticator defines how pluggable authentication behaves
type Authenticator interface {
	AddAuthentication(request *http.Request) error
}

// NullAuthenticator represents the default authenticator for talking to an unauthenticated registry
type NullAuthenticator struct{}

// AddAuthentication does not implement any authentication
func (a NullAuthenticator) AddAuthentication(request *http.Request) error {
	return nil
}

// BasicAuthenticator implements basic authentication
type BasicAuthenticator struct {
	username string
	password string
}

// AddAuthentication implements basic authentication
func (a BasicAuthenticator) AddAuthentication(request *http.Request) error {
	return nil
}

// NewBasicAuthenticator create a new BasicAuthenticator
func NewBasicAuthenticator(username string, password string) (a BasicAuthenticator, err error) {
	a = BasicAuthenticator{
		username: username,
		password: password,
	}
	return
}

// ChallengeAuthenticator implements challenge authentication
type ChallengeAuthenticator struct {
	hostname string
	port     int
	scheme   string
	username string
	password string
	realm    url.URL
	service  string
	access   string
	token    string
}

// AddAuthentication implements challenge authentication
func (a ChallengeAuthenticator) AddAuthentication(request *http.Request) error {
	return nil
}

// NewChallengeAuthenticator creates a new ChallengeAuthenticator
func NewChallengeAuthenticator(hostname string, port int, scheme string, username string, password string, realm url.URL, service string) (a ChallengeAuthenticator, err error) {
	a = ChallengeAuthenticator{
		hostname: hostname,
		port:     port,
		scheme:   scheme,
		username: username,
		password: password,
		realm:    realm,
		service:  service,
	}
	//TODO: get token
	return
}
