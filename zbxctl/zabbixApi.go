package zbxctl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/sgaunet/controls/results"
)

func New(cfg *ZbxCtl) (ZabbixApi, error) {
	z := ZabbixApi{
		cfg: *cfg,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
	data := zbxLogin{
		JsonRPC: "2.0",
		Method:  "user.login",
		Params: zbxParams{
			User:     z.cfg.User,
			Password: z.cfg.Password,
		},
		Id: 1,
	}

	postBody, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return z, err
	}
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", z.cfg.APIEndpoint, responseBody)
	if err != nil {
		return z, err
	}
	req = req.WithContext(context.TODO())
	req.Header.Set("Content-Type", "application/json")
	resp, err := z.client.Do(req)
	if err != nil {
		// log.Fatalf("An Error Occured %v", err)
		return z, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
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
	req, err := http.NewRequest("POST", z.cfg.APIEndpoint, responseBody)
	if err != nil {
		return err
	}
	req = req.WithContext(context.TODO())
	req.Header.Set("Content-Type", "application/json")
	resp, err := z.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
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
	since := strconv.FormatInt(time.Now().Add(time.Duration(-z.cfg.Since*int(time.Second))).Unix(), 10)

	dataGetProblem := zbxGetProblem{
		JsonRPC: "2.0",
		Method:  "problem.get",
		Params: zbxParamsProblem{
			Suppressed:   false,
			Recent:       false,
			Acknowledged: false,
			Time_from:    since,
			Tags:         z.cfg.FilterProblemsByTags,
		},
		Auth: z.auth,
		Id:   1,
	}

	enc, err := json.Marshal(dataGetProblem)
	if err != nil {
		return reportTable, err
	}
	req, err := http.NewRequest("POST", z.cfg.APIEndpoint, bytes.NewBuffer(enc))
	if err != nil {
		return reportTable, err
	}
	req = req.WithContext(context.TODO())
	req.Header.Set("Content-Type", "application/json")
	resp, err := z.client.Do(req)
	if err != nil {
		return reportTable, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return reportTable, err
	}
	defer resp.Body.Close()
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
			Result: fmt.Sprintf("Severity : %s (Threshold : %d)", pb.Severity, z.cfg.SeverityThreshold),
			Pass:   pbSeverity < z.cfg.SeverityThreshold,
		}
		reportTable = append(reportTable, r)
		r.PrintToStdout()
		idx++
	}
	return
}
