package main

import (
	"fmt"
	"io/ioutil"

	"sgaunet/controls/config"

	"gopkg.in/yaml.v2"
)

func ReadyamlConfigFile(filename string) (config.YamlConfig, error) {
	var yamlConfig config.YamlConfig

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

	return yamlConfig, err
}
