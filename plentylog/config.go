package plentylog

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	InternalProvider configProvider `yaml:"internalProvider"`
	FileFormat       format         `yaml:"fileFormat"`
}

type configProvider string

const (
	configProviderCLI  configProvider = "cli"
	configProviderFile configProvider = "file"
)

func loadConfig(file string) *config {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading YAML file: %v\n", err)

		return nil
	}

	var c config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Printf("error unmarshalling YAML file: %v\n", err)

		return nil
	}

	return &c
}
