package sshctl

import (
	"io/ioutil"
	"strings"
	"time"

	"sgaunet/controls/config"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func EscapeForMarkdown(str string) string {
	tmp := strings.ReplaceAll(str, "\n", "<br>")
	return strings.ReplaceAll(tmp, "|", "\\|")
}

func LaunchControls(cfgSrv []config.SSHServer, asserts []config.AssertSSH) [][]string {
	red := color.New(color.FgRed, color.Bold)
	resultTable := [][]string{{"Host", "Cmd", "Expected Result", "Result"}}

	idx := 0
	for _, server := range cfgSrv {
		var sshConfig ssh.ClientConfig

		if len(server.Sshkey) != 0 {
			auth := PublicKeyFile(server.Sshkey)
			if auth == nil {
				panic("Key not found")
			}
			sshConfig = ssh.ClientConfig{
				User: server.User,
				Auth: []ssh.AuthMethod{
					auth,
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				Timeout:         2 * time.Second,
			}
		} else {
			sshConfig = ssh.ClientConfig{
				User: server.User,
				Auth: []ssh.AuthMethod{
					ssh.Password(server.Password),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				Timeout:         2 * time.Second,
			}
		}

		connection, err := ssh.Dial("tcp", server.Host, &sshConfig)
		if err != nil {
			for _, control := range asserts {
				red.Printf("Host : %30s      -- Failed to connect\n", server.Host)
				resultTable = append(resultTable, []string{server.Host, EscapeForMarkdown(control.Cmd), "N/A", "<span style=\"color:red\">Failed to connect</span>"})
				idx++
			}
			continue
		}
		defer connection.Close()

		for _, control := range asserts {
			sshControl := sshControl{server: server.Host,
				cmd:            control.Cmd,
				expectedResult: control.ExpectedResult}

			var newOutput string
			session, err := connection.NewSession()
			if err != nil {
				sshControl.SetOutput("Failed to create session", false)
				red.Printf("Host : %30s      -- Failed to create session\n", server.Host)
				resultTable = append(resultTable, sshControl.GetResultLine())
				idx++
				continue
			}

			output, err := session.CombinedOutput(control.Cmd)
			if err != nil {
				sshControl.SetOutput("Failed to use session", false)
				red.Printf("Host : %30s      -- Failed to use session\n", server.Host)
				resultTable = append(resultTable, sshControl.GetResultLine())
				continue
			}
			if len(output) != 0 {
				newOutput = string(output)[0 : len(string(output))-1] // Remove the EOL
			}
			// session.Wait()
			session.Close()

			if control.ExpectedResult == newOutput {
				sshControl.SetOutput(newOutput, true)
				resultTable = append(resultTable, sshControl.GetResultLine())
				sshControl.PrintToStdout()
			} else {
				sshControl.SetOutput(newOutput, false)
				resultTable = append(resultTable, sshControl.GetResultLine())
				sshControl.PrintToStdout()
			}
			idx++
		}
	}

	return resultTable
}
