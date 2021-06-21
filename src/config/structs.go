package config

type DbConfig struct {
	Dbhost      string `yaml:"dbhost"`
	Dbport      string `yaml:"dbport"`
	Dbuser      string `yaml:"dbuser"`
	Dbpassword  string `yaml:"dbpassword"`
	Dbname      string `yaml:"dbname"`
	Dbsizelimit int    `yaml:"sizelimit"`
}

type Servers struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Sshkey   string `yaml:"sshkey"`
}

type AssertSSH struct {
	Cmd            string `yaml:"cmd"`
	ExpectedResult string `yaml:"expected"`
}

type AssertHTTP struct {
	Host       string `yaml:"host"`
	HostHeader string `yaml:"hostheader"`
	Comment    string `yaml:"comment"`
}

// Struct representing the yaml configuration file passed as a parameter to the program
type YamlConfig struct {
	Db          []DbConfig             `yaml:"db"`
	SshServers  map[string][]Servers   `yaml:"sshServers"`
	SshAsserts  map[string][]AssertSSH `yaml:"sshAsserts"`
	AssertsHTTP []AssertHTTP           `yaml:"assertsHTTP"`
	ZbxCtl      ZbxCtl                 `yaml:"zbxCtl"`
}

type ZbxCtl struct {
	ApiEndpoint       string `yaml:"apiEndpoint"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	Since             int    `yaml:"since"`
	SeverityThreshold int    `yaml:"severityThreshold"`
}
