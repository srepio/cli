package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrInvalidConnection = errors.New("config error: invalid current connection")
	ErrNoConfig          = errors.New("config error: config file not found")

	DefaultPath string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		DefaultPath = ".srep.yaml"
	} else {
		DefaultPath = fmt.Sprintf("%s/.srep.yaml", home)
	}
}

type Config struct {
	Connections       []Api  `yaml:"connections"`
	CurrentConnection string `yaml:"current_connection"`
}

type Api struct {
	Name   string `yaml:"name"`
	Url    string `yaml:"api_url"`
	Token  string `yaml:"token"`
	Scheme string `yaml:"scheme"`
}

// Load the config from ~/.srep.yaml or from the value of SREP_CONFIG
// Fails if the file doesn't exist
func GetConfig() (*Config, error) {
	filePath := getPath()

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

func getPath() string {
	filePath := DefaultPath
	if _, ok := os.LookupEnv("SREP_CONFIG"); ok {
		filePath = os.Getenv("SREP_CONFIG")
	}
	return filePath
}

func Initialise() error {
	filePath := getPath()

	c := &Config{
		Connections: []Api{
			{
				Name:   "default",
				Url:    "api.srep.io",
				Scheme: "https",
				Token:  "",
			},
		},
		CurrentConnection: "default",
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, b, 0600); err != nil {
		return err
	}
	return nil
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

func (c *Config) GetCurrentConnection() Api {
	for _, conn := range c.Connections {
		if conn.Name == c.CurrentConnection {
			return conn
		}
	}
	return Api{}
}

func (c *Config) SetConnection(conn Api) error {
	// Check if we should overwrite an existing context first
	for i := range c.Connections {
		if c.Connections[i].Name == conn.Name {
			c.Connections[i] = conn
			return nil
		}
	}
	c.Connections = append(c.Connections, conn)
	return c.Persist()
}

func (c *Config) SetContext(name string) error {
	for _, conn := range c.Connections {
		if conn.Name == name {
			c.CurrentConnection = name
			return c.Persist()
		}
	}
	return errors.New("unknown connection name")
}

// Persist the current config
func (c *Config) Persist() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(getPath(), data, 0600)
}
