package sshctl

import (
	"sgaunet/controls/config"
)

type Server struct {
	Cfg config.Servers
}

type sshControl struct {
	server         string
	cmd            string
	output         string
	expectedResult string
	passed         bool
}
