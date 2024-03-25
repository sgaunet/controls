package sshctl

import "github.com/sgaunet/controls/internal/sshserver"

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

// Struct representing an SSH assertion
type AssertSSH struct {
	Title          string `yaml:"title" validate:"required"`
	Cmd            string `yaml:"cmd" validate:"required"`
	ExpectedResult string `yaml:"expected" validate:"required"`
}
