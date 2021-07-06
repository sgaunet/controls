package sshctl

import "github.com/fatih/color"

func (c *sshControl) SetOutput(output string, passed bool) {
	c.passed = passed
	c.output = output
}

func (c *sshControl) GetResultLine() []string {
	var result string
	if c.passed {
		result = "<span style=\"color:green\">" + c.output + "</span>"
	} else {
		result = "<span style=\"color:red\">" + c.output + "</span>"
	}
	return []string{c.server, EscapeForMarkdown(c.cmd), EscapeForMarkdown(c.expectedResult), EscapeForMarkdown(result)}
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
