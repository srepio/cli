package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/srepio/sdk/types"
	"gopkg.in/yaml.v3"
)

var (
	ErrInvalidConnection = errors.New("config error: invalid current connection")
	ErrNoConfig          = errors.New("config error: config file not found")
)

type Config struct {
	DefaultDriver     types.DriverName `yaml:"default_driver"`
	Connections       []Api            `yaml:"connections"`
	CurrentConnection string           `yaml:"current_connection"`
}

type Api struct {
	Name  string `yaml:"name"`
	Url   string `yaml:"api_url"`
	Token string `yaml:"token"`
}

// Load the config from ~/.srep.yaml or from the value of SREP_CONFIG
// Fails if the file doesn't exist
func GetConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := fmt.Sprintf("%s/.srep.yaml", home)

	if _, ok := os.LookupEnv("SREP_CONFIG"); ok {
		filePath = os.Getenv("SREP_CONFIG")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrNoConfig
		}
		return nil, err
	}

	config := &Config{}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	if err := config.validate(); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) validate() error {
	// Validate current connection
	cv := false
	for _, conn := range c.Connections {
		if conn.Name == c.CurrentConnection {
			cv = true
		}
	}
	if !cv {
		return ErrInvalidConnection
	}

	return nil
}
