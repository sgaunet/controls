package config

import (
	"github.com/sgaunet/controls/httpctl"
	"github.com/sgaunet/controls/postgresctl"
	"github.com/sgaunet/controls/sshctl"
	"github.com/sgaunet/controls/sshserver"
	"github.com/sgaunet/controls/zbxctl"
)

// Struct representing the yaml configuration file passed as a parameter to the program
type YamlConfig struct {
	Db          []postgresctl.DbConfig           `yaml:"db"`
	SshServers  map[string][]sshserver.SSHServer `yaml:"sshServers"`
	SshAsserts  map[string][]sshctl.AssertSSH    `yaml:"sshAsserts"`
	AssertsHTTP []httpctl.AssertHTTP             `yaml:"assertsHTTP"`
	ZbxCtl      zbxctl.ZbxCtl                    `yaml:"zbxCtl"`
}
