package httpctl

// Struct representing an HTTP assertion
type AssertHTTP struct {
	Host       string `yaml:"host" validate:"required"`
	HostHeader string `yaml:"hostheader" `
	Title      string `yaml:"title" validate:"required"`
}
