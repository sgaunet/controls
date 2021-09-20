package sshctl

import "github.com/sgaunet/controls/sshserver"

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

//Struct representing an SSH assertion
type AssertSSH struct {
	Cmd            string `yaml:"cmd"`
	ExpectedResult string `yaml:"expected"`
}
