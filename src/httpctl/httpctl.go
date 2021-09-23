package httpctl

import (
	"net/http"

	"github.com/sgaunet/controls/results"
)

func ctlHTTP(assertHTTP AssertHTTP) results.Result {
	var r results.Result
	r.Title = assertHTTP.Title

	req, err := http.NewRequest("GET", assertHTTP.Host, nil)
	if err != nil {
		//c.statusCode = err.Error()
		r.Result = err.Error()
		return r
	}
	if len(assertHTTP.HostHeader) != 0 {
		req.Host = assertHTTP.HostHeader
	}
	// client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
	// 	return http.ErrUseLastResponse
	// }}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		r.Result = err.Error()
		//c.statusCode = err.Error()
		return r
	}
	// _, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error", err.Error())
	// }
	//fmt.Println(resp.StatusCode)
	// statusCode := strconv.Itoa(resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		r.Pass = true
	}

	return r
}

func LaunchControls(asserts []AssertHTTP) (resultTable []results.Result) {
	idx := 0
	for _, assert := range asserts {
		res := ctlHTTP(assert)
		res.PrintToStdout()
		resultTable = append(resultTable, res)
		idx++
	}
	return resultTable
}
