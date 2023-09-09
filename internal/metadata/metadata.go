package metadata

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed metadata.json
var mdJson string

type Metadata []Scenario

type Scenario struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty"`
	Version     string   `json:"version"`
	Tags        []string `json:"tags"`
	Ports       []Port   `json:"ports"`
	Volumes     []Volume `json:"volumes"`
	Privileged  bool     `json:"privileged"`
}

type Port struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

type Volume struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

func Get() (*Metadata, error) {
	md := &Metadata{}

	if err := json.Unmarshal([]byte(mdJson), md); err != nil {
		return nil, err
	}

	return md, nil
}

func Find(name string) (*Scenario, error) {
	md, err := Get()
	if err != nil {
		return nil, err
	}

	for _, s := range *md {
		if s.Name == name {
			return &s, nil
		}
	}

	return nil, fmt.Errorf("unknown senario %s", name)
}
