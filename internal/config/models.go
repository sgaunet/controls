package config

import (
	"github.com/sgaunet/controls/internal/httpctl"
	"github.com/sgaunet/controls/internal/postgresctl"
	"github.com/sgaunet/controls/internal/sshctl"
	"github.com/sgaunet/controls/internal/sshserver"
)

// Struct representing the yaml configuration file passed as a parameter to the program
type YamlConfig struct {
	Db          []postgresctl.DbConfig           `yaml:"db"`
	SshServers  map[string][]sshserver.SSHServer `yaml:"sshServers"`
	SshAsserts  map[string][]sshctl.AssertSSH    `yaml:"sshAsserts"`
	AssertsHTTP []httpctl.AssertHTTP             `yaml:"assertsHTTP" validate:"dive"`
	ZbxCtl      *ZbxCtlConfig                    `yaml:"zbxCtl"`
}

// Configuration of zabbix controls
type ZbxCtlConfig struct {
	APIEndpoint          string                       `yaml:"apiEndpoint" validate:"required"`
	User                 string                       `yaml:"user" validate:"required"`
	Password             string                       `yaml:"password" validate:"required"`
	Since                int                          `yaml:"since"`
	SeverityThreshold    int                          `yaml:"severityThreshold"`
	FilterProblemsByTags []ZbxTagsFilterProblemConfig `yaml:"filterProblemsByTags" validate:"dive"`
}

// Configuration of a tag filter for zabbix problems
// Operator can be:
// 0 - (default) Like;
// 1 - Equal;
// 2 - Not like;
// 3 - Not equal
// 4 - Exists;
// 5 - Not exists.
type ZbxTagsFilterProblemConfig struct {
	Tag      string `json:"tag" yaml:"tag" validate:"required"`
	Value    string `json:"value" yaml:"value"`
	Operator string `json:"operator" yaml:"operator" validate:"required,oneof='like' 'equal' 'not like' 'not equal' 'exists' 'not exists'"`
}
