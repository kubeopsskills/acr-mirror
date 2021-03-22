package acr

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Target struct {
	Name                    string  `yaml:"name"`
	Prefix                  *string `yaml:"prefix"`
	RePrefix                *string `yaml:"rePrefix"`
	TargetResourceGroupName string  `yaml:"resourceGroupName"`
}

type Registries struct {
	Name                    string   `yaml:"name"`
	Repositories            []string `yaml:"repositories"`
	SourceResourceGroupName string   `yaml:"resourceGroupName"`
	Tags                    []string `yaml:"tags"`
	Target                  []Target `yaml:"target"`
}

type Config struct {
	SourceSubscriptionID string       `yaml:"source_subscription_id"`
	TargetSubscriptionID string       `yaml:"target_subscription_id"`
	Registries           []Registries `yaml:"registries"`
}

func (config *Config) GetConfig(configFileName string) *Config {
	configYAMLFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Fatalf("Error: #%v", err)
	}
	err = yaml.Unmarshal(configYAMLFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}
