package config

import (
	"github.com/sgaunet/controls/internal/httpctl"
	"github.com/sgaunet/controls/internal/postgresctl"
	"github.com/sgaunet/controls/internal/sshctl"
	"github.com/sgaunet/controls/internal/sshserver"
	"github.com/sgaunet/controls/internal/zbxctl"
)

// Struct representing the yaml configuration file passed as a parameter to the program
type YamlConfig struct {
	Db          []postgresctl.DbConfig           `yaml:"db"`
	SshServers  map[string][]sshserver.SSHServer `yaml:"sshServers"`
	SshAsserts  map[string][]sshctl.AssertSSH    `yaml:"sshAsserts"`
	AssertsHTTP []httpctl.AssertHTTP             `yaml:"assertsHTTP"`
	ZbxCtl      zbxctl.ZbxCtl                    `yaml:"zbxCtl"`
}
