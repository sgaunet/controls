package sshctl

import "sgaunet/controls/sshserver"

type Server struct {
	Cfg sshserver.SSHServer
}

type sshControl struct {
	server         string
	cmd            string
	output         string
	expectedResult string
	passed         bool
}
