package registry

type Authenticator interface {
	Test() error
}