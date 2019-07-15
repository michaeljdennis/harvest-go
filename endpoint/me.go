package endpoint

// Me ...
type Me struct {
	*Endpoint
}

// NewMe returns new Me
func NewMe() *Me {
	return &Me{
		Endpoint: NewEndpoint(
			"users/me",
		),
	}
}
