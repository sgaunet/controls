package config

// Struct representing the yaml configuration file passed as a parameter to the program
type YamlConfig struct {
	Db          []DbConfig             `yaml:"db"`
	SshServers  map[string][]Servers   `yaml:"sshServers"`
	SshAsserts  map[string][]AssertSSH `yaml:"sshAsserts"`
	AssertsHTTP []AssertHTTP           `yaml:"assertsHTTP"`
	ZbxCtl      ZbxCtl                 `yaml:"zbxCtl"`
}

// Struct to represent the informations to connect to the Zabbix API
type ZbxCtl struct {
	ApiEndpoint       string `yaml:"apiEndpoint"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	Since             int    `yaml:"since"`
	SeverityThreshold int    `yaml:"severityThreshold"`
}

// Struct representing the connection to a postgres Database
type DbConfig struct {
	Dbhost      string `yaml:"dbhost"`
	Dbport      string `yaml:"dbport"`
	Dbuser      string `yaml:"dbuser"`
	Dbpassword  string `yaml:"dbpassword"`
	Dbname      string `yaml:"dbname"`
	Dbsizelimit int    `yaml:"sizelimit"`
}

// Struct representing the connection to an SSH Server
type Servers struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Sshkey   string `yaml:"sshkey"`
}

//Struct representing an SSH assertion
type AssertSSH struct {
	Cmd            string `yaml:"cmd"`
	ExpectedResult string `yaml:"expected"`
}

//Struct representing an HTTP assertion
type AssertHTTP struct {
	Host       string `yaml:"host"`
	HostHeader string `yaml:"hostheader"`
	Comment    string `yaml:"comment"`
}
