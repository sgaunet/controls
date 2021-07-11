package config

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v2"
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

func ReadyamlConfigFile(filename string) (YamlConfig, error) {
	var yamlConfig YamlConfig

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return yamlConfig, err
	}

	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		return yamlConfig, err
	}

	// Add default ports to SSH Servers
	for idxGrp := range yamlConfig.SshServers {
		for idxSrv := range yamlConfig.SshServers[idxGrp] {
			yamlConfig.SshServers[idxGrp][idxSrv].AddPortIfNotPresent()
		}
	}
	return yamlConfig, err
}
