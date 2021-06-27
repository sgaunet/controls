package httpctl

type httpControl struct {
	url        string
	hostHeader string
	statusCode string
	comment    string
	passed     bool
}
