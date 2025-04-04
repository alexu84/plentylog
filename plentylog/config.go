package plentylog

import (
	"os"

	"gopkg.in/yaml.v2"
)

// config represents the configuration structure for the logger
// It contains the internal provider and file format.
type config struct {
	InternalProvider configProvider `yaml:"internalProvider"`
	FileFormat       format         `yaml:"fileFormat"`
}

// configProvider represents the type of provider used for logging
type configProvider string

const (
	// configProviderCLI represents the CLI provider
	configProviderCLI configProvider = "cli"
	// configProviderFile represents the file provider
	configProviderFile configProvider = "file"
)

// loadConfig loads the configuration from a YAML file
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
