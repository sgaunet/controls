package sshctl

import (
	"io/ioutil"
	"strings"
	"time"

	"sgaunet/controls/config"

	"github.com/atsushinee/go-markdown-generator/doc"
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

func LaunchControls(cfgSrv []config.Servers, asserts []config.AssertSSH, report *doc.MarkDownDoc) *doc.MarkDownDoc {
	red := color.New(color.FgRed, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	t := doc.NewTable(len(cfgSrv)*len(asserts), 4)
	t.SetTitle(0, "Host")
	t.SetTitle(1, "Cmd")
	t.SetTitle(2, "Result Expected")
	t.SetTitle(3, "Result")

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

		connection, err := ssh.Dial("tcp", server.Host+":22", &sshConfig)
		if err != nil {
			for _, control := range asserts {
				// 				red.Printf("%-30s | %-60s | %-30s | %-30s |\n", server.Host, control.Cmd, "N/A", "Failed to connect")
				red.Printf("Host : %30s      -- Failed to connect\n", server.Host)
				t.SetContent(idx, 0, server.Host)
				t.SetContent(idx, 1, EscapeForMarkdown(control.Cmd))
				t.SetContent(idx, 2, "N/A")
				t.SetContent(idx, 3, "<span style=\"color:red\">Failed to connect</span>")
				idx++
			}
			continue
		}
		defer connection.Close()

		for _, control := range asserts {
			var newOutput string
			session, err := connection.NewSession()
			if err != nil {
				// fmt.Printf("Failed to create session: %s\n", err)
				red.Printf("Host : %30s      -- Failed to create session\n", server.Host)
				t.SetContent(idx, 0, server.Host)
				t.SetContent(idx, 1, EscapeForMarkdown(control.Cmd))
				t.SetContent(idx, 2, EscapeForMarkdown(control.ExpectedResult))
				t.SetContent(idx, 3, "<span style=\"color:red\">Failed to create session</span>")
				idx++
				continue
			}

			output, err := session.CombinedOutput(control.Cmd)
			if err != nil {
				// fmt.Printf("Failed to use session: %s\n", err)
				red.Printf("Host : %30s      -- Failed to use session\n", server.Host)
				t.SetContent(idx, 0, server.Host)
				t.SetContent(idx, 1, EscapeForMarkdown(control.Cmd))
				t.SetContent(idx, 2, EscapeForMarkdown(control.ExpectedResult))
				t.SetContent(idx, 3, "<span style=\"color:red\">Failed to use session</span>")
				//idx++
				continue
			}
			if len(output) != 0 {
				newOutput = string(output)[0 : len(string(output))-1] // Remove the EOL
			}
			// session.Wait()
			session.Close()

			if control.ExpectedResult == newOutput {
				//green.Printf("%-30s | %-60s | %-30s | %-30s |\n", server.Host, control.Cmd, control.ExpectedResult, newOutput)
				green.Printf("Host            : %s\n", server.Host)
				green.Printf("Command         : %s\n", control.Cmd)
				green.Printf("Expected Result : %s\n", control.ExpectedResult)
				green.Printf("Output          : %s\n\n", newOutput)
				t.SetContent(idx, 0, server.Host)
				t.SetContent(idx, 1, EscapeForMarkdown(control.Cmd))
				t.SetContent(idx, 2, EscapeForMarkdown(control.ExpectedResult))
				t.SetContent(idx, 3, EscapeForMarkdown("<span style=\"color:green\">"+newOutput+"</span>"))
			} else {
				//red.Printf("%-30s | %-60s | %-30s | %-30s |\n", server.Host, control.Cmd, control.ExpectedResult, newOutput)
				red.Printf("Host            : %s\n", server.Host)
				red.Printf("Command         : %s\n", control.Cmd)
				red.Printf("Expected Result : %s\n", control.ExpectedResult)
				red.Printf("Output          : %s\n\n", newOutput)
				t.SetContent(idx, 0, server.Host)
				t.SetContent(idx, 1, EscapeForMarkdown(control.Cmd))
				t.SetContent(idx, 2, EscapeForMarkdown(control.ExpectedResult))
				t.SetContent(idx, 3, EscapeForMarkdown("<span style=\"color:red\">"+newOutput+"</span>"))
			}
			idx++
		}
		//idx++
	}

	report.WriteTable(t)
	report.WriteLines(1)
	return report
}
