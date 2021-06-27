package sshctl

import "github.com/fatih/color"

func (c *sshControl) SetOutput(output string, passed bool) {
	if passed {
		c.output = "<span style=\"color:green\">" + output + "</span>"
	} else {
		c.output = "<span style=\"color:red\">" + output + "</span>"
	}
}

func (c *sshControl) GetResultLine() []string {
	return []string{c.server, EscapeForMarkdown(c.cmd), EscapeForMarkdown(c.expectedResult), EscapeForMarkdown(c.output)}
}

func (c *sshControl) PrintToStdout() {
	var std *color.Color
	if c.passed {
		std = color.New(color.FgGreen, color.Bold)
	} else {
		std = color.New(color.FgRed, color.Bold)
	}

	std.Printf("Host            : %s\n", c.server)
	std.Printf("Command         : %s\n", c.cmd)
	std.Printf("Expected Result : %s\n", c.expectedResult)
	std.Printf("Output          : %s\n\n", c.output)
}
