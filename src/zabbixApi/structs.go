package zabbixapi

type ZabbixApi struct {
	auth              string
	url               string
	since             int
	severityThreshold int
}

type zbxParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type zbxLogin struct {
	JsonRPC string    `json:"jsonrpc"`
	Method  string    `json:"method"`
	Params  zbxParams `json:"params"`
	Id      int       `json:"id"`
}

type zbxLogout struct {
	JsonRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
	Auth    string            `json:"auth"`
	Id      int               `json:"id"`
}

type zbxLoginReturn struct {
	JsonRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      int    `json:"id"`
}

type zbxParamsProblem struct {
	Suppressed   bool   `json:"suppressed"`
	Recent       bool   `json:"recent"`
	Acknowledged bool   `json:"acknowledged"`
	Time_from    string `json:"time_from"`
}

type zbxGetProblem struct {
	JsonRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  zbxParamsProblem `json:"params"`
	Auth    string           `json:"auth"`
	Id      int              `json:"id"`
}

type zbxProblem struct {
	Acknowledged  string `json:"acknowledged"`
	Clock         string `json:"clock"`
	Correlationid string `json:"correlationid"`
	Eventid       string `json:"eventid"`
	Name          string `json:"name"`
	Ns            string `json:"ns"`
	Object        string `json:"object"`
	Objectid      string `json:"objectid"`
	Opdata        string `json:"opdata"`
	R_clock       string `json:"r_clock"`
	R_eventid     string `json:"r_eventid"`
	R_ns          string `json:"r_ns"`
	Severity      string `json:"severity"`
	Source        string `json:"source"`
	Suppressed    string `json:"suppressed"`
}

type zbxResultProblem struct {
	JsonRPC string       `json:"jsonrpc"`
	Result  []zbxProblem `json:"result"`
	Id      int          `json:"id"`
}
