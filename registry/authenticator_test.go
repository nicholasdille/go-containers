package registry

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestNewNullAuthenticator(t *testing.T) {
	a, err := NewNullAuthenticator()
	if err != nil {
		t.Errorf("Failed to create new NullAuthenticator")
	}
	if reflect.TypeOf(a).String() != "registry.NullAuthenticator" {
		t.Errorf("Failed to create new NullAuthenticator: Unexpected type <%s>", reflect.TypeOf(a))
	}
}

func TestAddNullAuthentication(t *testing.T) {
	a := NullAuthenticator{}

	r, err := http.NewRequest("GET", "http://localhost", strings.NewReader(""))
	if err != nil {
		t.Errorf("Failed to create new request")
	}

	err = a.AddAuthentication(r)
	if err != nil {
		t.Errorf("NullAuthenticator failed to AddAuthentication")
	}
}

func TestNewBasicAuthenticator(t *testing.T) {
	a, err := NewBasicAuthenticator("user", "pass")
	if err != nil {
		t.Errorf("Failed to create new BasicAuthenticator")
	}

	if reflect.TypeOf(a).String() != "registry.BasicAuthenticator" {
		t.Errorf("Failed to create new BasicAuthenticator: Unexpected type <%s>", reflect.TypeOf(a))
	}

	if a.username != "user" {
		t.Errorf("Failed to set username on BasicAuthenticator: Expected <user> but got <%s>", a.username)
	}
	if a.password != "pass" {
		t.Errorf("Failed to set password on BasicAuthenticator: Expected <pass> but got <%s>", a.password)
	}
}

func TestAddBasicAuthentication(t *testing.T) {
	a, err := NewBasicAuthenticator("user", "pass")
	if err != nil {
		t.Errorf("Failed to create new BasicAuthenticator")
	}

	r, err := http.NewRequest("GET", "http://localhost", strings.NewReader(""))
	if err != nil {
		t.Errorf("Failed to create new request")
	}

	err = a.AddAuthentication(r)
	if err != nil {
		t.Errorf("NullAuthenticator failed to AddAuthentication")
	}

	username, password, ok := r.BasicAuth()
	if ! ok {
		t.Errorf("Failed to retrieve basic authentication")
	}
	if username != a.username {
		t.Errorf("Username does not match: Expected <%s> but got <%s>", a.username, username)
	}
	if password != a.password {
		t.Errorf("Password does not match: Expected <%s> but got <%s>", a.password, password)
	}
}
