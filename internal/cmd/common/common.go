package common

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/config"
	"github.com/srepio/sdk/client"
)

var (
	srep *client.Client
)

func ScenarioFlags(cmd *cobra.Command) {
	cmd.Flags().String("driver", "docker", "Determine which drive to use. Default: docker.")
}

func ScenarioCompletion() []string {
	out := []string{}
	s, err := Client().Getscenarios(context.Background())
	if err != nil {
		return out
	}

	for _, scenario := range *s.Scenarios {
		out = append(out, scenario.Name)
	}

	return out
}

func InitClient(config *config.Config) {
	conn := config.GetCurrentConnection()

	srep = client.NewClient(&client.ClientOptions{
		Url:    conn.Url,
		Token:  conn.Token,
		Scheme: conn.Scheme,
	})
}

func Client() *client.Client {
	if srep == nil {
		srep = client.NewClient(&client.ClientOptions{})
	}
	return srep
}
