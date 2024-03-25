package sshserver

import (
	"regexp"
	"strconv"
)

func (s *SSHServer) IsSSHPortDefined() bool {
	re := regexp.MustCompile(".*:([0-9]+)$")
	match := re.FindStringSubmatch(s.Host)
	if len(match) == 2 {
		_, err := strconv.Atoi(match[1])
		return err == nil
	}

	return false
}

func (s *SSHServer) AddPortIfNotPresent() {
	if !s.IsSSHPortDefined() {
		s.Host = s.Host + ":22"
	}
}
