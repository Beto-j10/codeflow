package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// CompilerConfig is the config for the compiler
type CompilerConfig struct {
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	URL          string `yaml:"URL"`
}

// DgraphConfig is the config for the dgraph
type DgraphConfig struct {
	Port string `yaml:"port"`
}

// ServerConfig is the config for the server
type ServerConfig struct {
	Port string `yaml:"port"`
}

// Config contains the configuration for the server
type Config struct {
	Compiler CompilerConfig `yaml:"Compiler"`
	Dgraph   DgraphConfig   `yaml:"Dgraph"`
	Server   ServerConfig   `yaml:"Server"`
}

// LoadConfig loads the config file
func LoadConfig(filename string) (*Config, error) {
	config := &Config{}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
