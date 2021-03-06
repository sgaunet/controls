package zbxctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/sgaunet/controls/results"
)

func New(login string, password string, url string, sinceMs int, severityThreshold int) (ZabbixApi, error) {
	var z ZabbixApi
	z.url = url
	z.since = sinceMs
	z.severityThreshold = severityThreshold
	data := zbxLogin{
		JsonRPC: "2.0",
		Method:  "user.login",
		Params: zbxParams{
			User:     login,
			Password: password,
		},
		Id: 1,
	}

	postBody, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return z, err
	}
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		// log.Fatalf("An Error Occured %v", err)
		return z, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return z, err
	}
	var zbxResp zbxLoginReturn
	err = json.Unmarshal(body, &zbxResp)
	if err != nil {
		fmt.Printf("Error : %s %d %s\n", zbxResp.Result, resp.StatusCode, "cannot login to zabbix API")
		return z, errors.New("cannot login to zabbix API")
	}
	if len(zbxResp.Result) == 0 {
		return z, errors.New("cannot login to zabbix API")
	}
	z.auth = zbxResp.Result
	return z, err
}

func (z *ZabbixApi) Logout() error {
	data2 := zbxLogout{
		JsonRPC: "2.0",
		Method:  "user.logout",
		Params:  make(map[string]string),
		Id:      1,
		Auth:    z.auth,
	}

	postBody, _ := json.Marshal(data2)
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(z.url, "application/json", responseBody)
	return err
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
}

func (z *ZabbixApi) FailedResultControls(err error) (reportTable []results.Result) {
	reportTable = append(reportTable, results.Result{
		Title:  "API Access",
		Result: err.Error(),
		Pass:   false,
	})
	//reportTable = append(reportTable, []string{z.url, "<span style=\"color:red\">" + err.Error() + "</span>"})
	return
}

func (z *ZabbixApi) LaunchControls() (reportTable []results.Result, err error) {
	since := strconv.FormatInt(time.Now().Add(time.Duration(-z.since*int(time.Second))).Unix(), 10)

	dataGetProblem := zbxGetProblem{
		JsonRPC: "2.0",
		Method:  "problem.get",
		Params: zbxParamsProblem{
			Suppressed:   false,
			Recent:       false,
			Acknowledged: false,
			Time_from:    since,
		},
		Auth: z.auth,
		Id:   1,
	}

	enc, _ := json.Marshal(dataGetProblem)
	resp, err := http.Post(z.url, "application/json", bytes.NewBuffer(enc))
	//Handle Error
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var obj map[string]interface{}
	json.Unmarshal(body, &obj)

	var resultProblems zbxResultProblem
	json.Unmarshal(body, &resultProblems)

	// reportTable = append(reportTable, []string{"Problem", "Severity", "Result"})
	idx := 0

	for _, pb := range resultProblems.Result {
		pbSeverity, err := strconv.Atoi(pb.Severity)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		r := results.Result{
			Title:  pb.Name,
			Result: fmt.Sprintf("Severity : %s (Threshold : %d)", pb.Severity, z.severityThreshold),
			Pass:   pbSeverity < z.severityThreshold,
		}
		reportTable = append(reportTable, r)
		r.PrintToStdout()
		idx++
	}
	return
}
