package sshctl

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/sgaunet/controls/results"
	"github.com/sgaunet/controls/sshserver"

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

func LaunchControls(cfgSrv []sshserver.SSHServer, asserts []AssertSSH) (resultTable []results.Result) {
	red := color.New(color.FgRed, color.Bold)
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
				resultTable = append(resultTable, results.Result{
					Title:  control.Title,
					Result: "Failed to connect",
					Pass:   false,
				})
				idx++
			}
			continue
		}
		defer connection.Close()

		for _, control := range asserts {
			// sshControl := sshControl{server: server.Host,
			// 	cmd:            control.Cmd,
			// 	expectedResult: control.ExpectedResult}

			var newOutput string
			session, err := connection.NewSession()
			if err != nil {
				// sshControl.SetOutput("Failed to create session", false)
				// red.Printf("Host : %30s      -- Failed to create session\n", server.Host)
				r := results.Result{
					Title:  fmt.Sprintf("%s (%s)", control.Title, server.Host),
					Pass:   false,
					Result: "Failed to create SSH session",
				}
				resultTable = append(resultTable, r)
				r.PrintToStdout()
				idx++
				continue
			}

			output, err := session.CombinedOutput(control.Cmd)
			if err != nil {
				r := results.Result{
					Title:  fmt.Sprintf("%s (%s)", control.Title, server.Host),
					Pass:   false,
					Result: "Failed to create SSH session",
				}
				resultTable = append(resultTable, r)
				r.PrintToStdout()
				continue
			}
			if len(output) != 0 {
				newOutput = string(output)[0 : len(string(output))-1] // Remove the EOL
			}
			// session.Wait()
			session.Close()

			r := results.Result{
				Title:  fmt.Sprintf("%s (%s)", control.Title, server.Host),
				Result: newOutput,
				Pass:   control.ExpectedResult == newOutput,
			}
			if control.ExpectedResult != newOutput {
				r.Result = fmt.Sprintf("Result : %s   Expected: %s", newOutput, control.ExpectedResult)
			}
			r.PrintToStdout()
			resultTable = append(resultTable, r)
			idx++
		}
	}

	return resultTable
}
