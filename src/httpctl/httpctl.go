package httpctl

import (
	"net/http"
	"strconv"

	"sgaunet/controls/config"

	"github.com/fatih/color"
)

func (c *httpControl) GetResultLine() []string {
	var statusCode string
	if c.passed {
		statusCode = "<span style=\"color:green\">" + c.statusCode + "</span>"
	} else {
		statusCode = "<span style=\"color:red\">" + c.statusCode + "</span>"
	}
	return []string{c.url, statusCode, c.comment}
}

func (c *httpControl) PrintToStdout() {
	var std *color.Color
	if c.passed {
		std = color.New(color.FgGreen, color.Bold)
	} else {
		std = color.New(color.FgRed, color.Bold)
	}

	std.Printf("URL        : %s\n", c.url)
	std.Printf("StatusCode : %s\n", c.statusCode)
	std.Printf("Comment    : %s\n", c.comment)
}

func (c *httpControl) ctlHTTP(assertHTTP config.AssertHTTP) []string {
	req, err := http.NewRequest("GET", assertHTTP.Host, nil)
	if err != nil {
		c.statusCode = err.Error()
		return c.GetResultLine()
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
		c.statusCode = err.Error()
		return c.GetResultLine()
	}
	// _, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error", err.Error())
	// }
	//fmt.Println(resp.StatusCode)
	c.statusCode = strconv.Itoa(resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		c.passed = true
	} else {
		c.passed = false
	}
	return c.GetResultLine()
}

func LaunchControls(asserts []config.AssertHTTP) [][]string {
	//var urlToControl string
	// red := color.New(color.FgRed, color.Bold)
	// green := color.New(color.FgGreen, color.Bold)
	// nbAsserts := len(asserts)
	// t := doc.NewTable(nbAsserts, 5)
	// t.SetTitle(0, "URL")
	// t.SetTitle(1, "HTTP Status Code")
	// t.SetTitle(2, "Comment")
	resultTable := [][]string{{"URL", "Http StatusCode", "Comment"}}

	idx := 0
	for _, assert := range asserts {
		newHttpControl := httpControl{
			hostHeader: assert.HostHeader,
			url:        assert.Host,
			comment:    assert.Comment,
		}
		resultTable = append(resultTable, newHttpControl.ctlHTTP(assert))
		newHttpControl.PrintToStdout()

		// if err != nil {
		// 	red.Printf("URL: %s\n", urlToControl)
		// 	red.Printf("Error: %s\n", err.Error())
		// 	red.Printf("Comment: %s \n", assert.Comment)
		// 	// t.SetContent(idx, 0, urlToControl)
		// 	// t.SetContent(idx, 1, "<span style=\"color:red\">ERR :"+err.Error()+"</span>")
		// 	// t.SetContent(idx, 2, assert.Comment)
		// 	resultTable = append(resultTable, []string{urlToControl, "<span style=\"color:red\">ERR :" + err.Error() + "</span>", assert.Comment})
		// } else {
		// 	if !(statusCode >= 400) {
		// 		green.Printf("URL: %s\n", urlToControl)
		// 		green.Printf("StatusCode: %d\n", statusCode)
		// 		green.Printf("Comment: %s \n", assert.Comment)
		// 		//green.Printf("%-100s | %-20d | %-20s |\n", urlToControl, statusCode, assert.Comment)
		// 		resultTable = append(resultTable, []string{urlToControl, "<span style=\"color:green\">" + strconv.Itoa(statusCode) + "</span>", assert.Comment})
		// 	} else {
		// 		//red.Printf("%-100s | %-20d | %-20s |\n", urlToControl, statusCode, assert.Comment)
		// 		// t.SetContent(idx, 0, urlToControl)
		// 		// t.SetContent(idx, 1, "<span style=\"color:red\">"+strconv.Itoa(statusCode)+"</span>")
		// 		// t.SetContent(idx, 2, assert.Comment)
		// 		red.Printf("URL: %s\n", urlToControl)
		// 		red.Printf("StatusCode: %s\n", err.Error())
		// 		red.Printf("Comment: %s \n", assert.Comment)
		// 		//red.Printf("%-100s | %-20d | %-20s |\n", urlToControl, statusCode, assert.Comment)
		// 		resultTable = append(resultTable, []string{urlToControl, "<span style=\"color:red\">" + strconv.Itoa(statusCode) + "</span>", assert.Comment})
		// 	}
		// }
		idx++
	}

	//report.WriteTable(t)
	return resultTable
}
