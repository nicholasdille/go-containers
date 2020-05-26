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

func (r *Registry) Check() (test string, err error) {
	requester, err := NewRequester(r.hostname, r.port, r.insecure, r.username, r.password)
	if err != nil {
		err = errors.New("Failed to instantiate requester")
		return
	}

	test = requester.hostname

	return
}