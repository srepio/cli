package common

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/driver/docker"
	"github.com/srepio/sdk/client"
	"github.com/srepio/sdk/types"
)

var (
	srep *client.Client
)

func ScenarioFlags(cmd *cobra.Command) {
	cmd.Flags().String("driver", "docker", "Determine which drive to use. Default: docker.")
}

func ScenarioCompletion() []string {
	out := []string{}
	s, err := Client().GetMetadata(context.Background())
	if err != nil {
		return out
	}

	for _, scenario := range *s.Scenarios {
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

func Client() *client.Client {
	if srep == nil {
		srep = client.NewClient(&client.ClientOptions{})
	}
	return srep
}
