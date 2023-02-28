package cadence

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type CadenceConfig struct {
	Domain   string `yaml:"domain"`
	Service  string `yaml:"service"`
	HostPort string `yaml:"host"`
}

func SetupConfig(configPath string) *CadenceConfig {
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to log config file: %v, Error: %v", configPath, err))
	}

	config := &CadenceConfig{}

	if err := yaml.Unmarshal(configData, &config); err != nil {
		panic(fmt.Sprintf("Error initializing configuration: %v", err))
	}

	return config
}