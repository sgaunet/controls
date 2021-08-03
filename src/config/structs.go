package config

import (
	"sgaunet/controls/httpctl"
	"sgaunet/controls/postgresctl"
	"sgaunet/controls/sshctl"
	"sgaunet/controls/sshserver"
	"sgaunet/controls/zbxctl"
)

// Struct representing the yaml configuration file passed as a parameter to the program
type YamlConfig struct {
	Db          []postgresctl.DbConfig           `yaml:"db"`
	SshServers  map[string][]sshserver.SSHServer `yaml:"sshServers"`
	SshAsserts  map[string][]sshctl.AssertSSH    `yaml:"sshAsserts"`
	AssertsHTTP []httpctl.AssertHTTP             `yaml:"assertsHTTP"`
	ZbxCtl      zbxctl.ZbxCtl                    `yaml:"zbxCtl"`
}
