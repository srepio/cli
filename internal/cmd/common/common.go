package common

import (
	"errors"

	"github.com/srepio/cli/internal/driver/docker"
	"github.com/srepio/cli/internal/metadata"
	"github.com/srepio/sdk/types"
)

func ScenarioCompletion() []string {
	out := []string{}
	s, err := metadata.Get()

	if err != nil {
		return out
	}

	for _, scenario := range *s {
		out = append(out, scenario.Name)
	}

	return out
}

func GetDriver(k8s bool) (types.Driver, error) {
	if k8s {
		return nil, errors.New("k8s driver not implemented yet")
	}
	return docker.NewDockerDriver()
}
