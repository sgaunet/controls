package httpctl

//Struct representing an HTTP assertion
type AssertHTTP struct {
	Host       string `yaml:"host"`
	HostHeader string `yaml:"hostheader"`
	Title      string `yaml:"title"`
}
