package config

import (
	"sgaunet/controls/sshserver"
	"sgaunet/controls/zbxctl"
)

// Struct representing the yaml configuration file passed as a parameter to the program
type YamlConfig struct {
	Db          []DbConfig                       `yaml:"db"`
	SshServers  map[string][]sshserver.SSHServer `yaml:"sshServers"`
	SshAsserts  map[string][]AssertSSH           `yaml:"sshAsserts"`
	AssertsHTTP []AssertHTTP                     `yaml:"assertsHTTP"`
	ZbxCtl      zbxctl.ZbxCtl                    `yaml:"zbxCtl"`
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
