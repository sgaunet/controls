package httpctl

import (
	"fmt"
	"net/http"
	"strconv"

	"sgaunet/controls/config"

	"github.com/atsushinee/go-markdown-generator/doc"
	"github.com/fatih/color"
)

func ctlHTTP(assertHTTP config.AssertHTTP) (int, error) {
	req, err := http.NewRequest("GET", assertHTTP.Host, nil)
	if err != nil {
		fmt.Println("error", err.Error())
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
		return 0, err
	}
	// _, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error", err.Error())
	// }
	//fmt.Println(resp.StatusCode)
	return resp.StatusCode, err
}

func LaunchControls(asserts []config.AssertHTTP, report *doc.MarkDownDoc) *doc.MarkDownDoc {
	var urlToControl string
	red := color.New(color.FgRed, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	nbAsserts := len(asserts)
	t := doc.NewTable(nbAsserts, 5)
	t.SetTitle(0, "URL")
	t.SetTitle(1, "HTTP Status Code")
	t.SetTitle(2, "Comment")

	fmt.Printf("%-100s | %-20s | %-20s |\n", "URL", "HTTP Status code", "Comment")
	idx := 0
	for _, assert := range asserts {
		statusCode, err := ctlHTTP(assert)

		if len(assert.HostHeader) != 0 {
			urlToControl = assert.HostHeader
		} else {
			urlToControl = assert.Host
		}
		if err != nil {
			red.Printf("%-100s | %-20s | %-20s |\n", urlToControl, err.Error(), assert.Comment)
			t.SetContent(idx, 0, urlToControl)
			t.SetContent(idx, 1, "<span style=\"color:red\">ERR :"+err.Error()+"</span>")
			t.SetContent(idx, 2, assert.Comment)
		} else {
			if !(statusCode >= 400) {
				green.Printf("%-100s | %-20d | %-20s |\n", urlToControl, statusCode, assert.Comment)
				t.SetContent(idx, 0, urlToControl)
				t.SetContent(idx, 1, "<span style=\"color:green\">"+strconv.Itoa(statusCode)+"</span>")
				t.SetContent(idx, 2, assert.Comment)
			} else {
				red.Printf("%-100s | %-20d | %-20s |\n", urlToControl, statusCode, assert.Comment)
				t.SetContent(idx, 0, urlToControl)
				t.SetContent(idx, 1, "<span style=\"color:red\">"+strconv.Itoa(statusCode)+"</span>")
				t.SetContent(idx, 2, assert.Comment)
			}
		}
		idx++
	}

	report.WriteTable(t)
	return report
}
