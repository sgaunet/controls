package sshctl

import (
	"sgaunet/controls/config"
)

type Server struct {
	Cfg config.SSHServer
}

type sshControl struct {
	server         string
	cmd            string
	output         string
	expectedResult string
	passed         bool
}
