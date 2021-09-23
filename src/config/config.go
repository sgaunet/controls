package config

import (
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

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

func (y *YamlConfig) CheckAsserts() (confOK bool) {
	var red *color.Color
	red = color.New(color.FgRed, color.Bold)

	confOK = true
	for _, asserts := range y.SshAsserts {
		for _, assert := range asserts {
			if assert.Title == "" {
				confOK = false
				red.Printf("An SSH assert has no title : %v\n", assert)
			}
		}
	}

	for _, assert := range y.AssertsHTTP {
		if assert.Title == "" {
			confOK = false
			red.Printf("An HTTP assert has no title : %v\n", assert)
		}
	}
	return confOK
}
