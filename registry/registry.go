package registry

type Registry struct {
	hostname string
	port int
	insecure bool
	username string
	password string
}

func (r *Registry) Connect() (err error) {
	return nil
}

func NewRegistry() Registry {
	return Registry{
		hostname: "localhost",
		port: 5000,
		insecure: true,
		username: "",
		password: "",
	}
}