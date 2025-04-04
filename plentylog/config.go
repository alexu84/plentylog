package plentylog

import (
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

func loadConfig(file string) (*config, error) {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var c config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
