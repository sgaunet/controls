package httpctl

type httpControl struct {
	url        string
	hostHeader string
	statusCode string
	passed     bool
}

//Struct representing an HTTP assertion
type AssertHTTP struct {
	Host       string `yaml:"host"`
	HostHeader string `yaml:"hostheader"`
	Comment    string `yaml:"comment"`
}
