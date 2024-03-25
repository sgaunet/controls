package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

func ReadyamlConfigFile(filename string) (YamlConfig, error) {
	var yamlConfig YamlConfig

	yamlFile, err := os.ReadFile(filename)
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

func (y *YamlConfig) IsValid() bool {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(y)
	return err == nil
}

func (y *YamlConfig) Errors() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(y)
	return err
}
