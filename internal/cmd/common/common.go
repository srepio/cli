package common

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/driver/docker"
	"github.com/srepio/cli/internal/metadata"
	"github.com/srepio/sdk/types"
)

func ScenarioFlags(cmd *cobra.Command) {
	cmd.Flags().String("driver", "docker", "Determine which drive to use. Default: docker.")
}

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

func GetDriver(driver string) (types.Driver, error) {
	switch driver {
	case "docker":
		return docker.NewDockerDriver()
	default:
		return nil, fmt.Errorf("unknwon driver %s", driver)
	}
}
