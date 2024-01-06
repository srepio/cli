/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package run

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/sdk/client"
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "run [scenario]",
		Short:     "Run the specified practice scenarios",
		GroupID:   "srep",
		Args:      cobra.ExactArgs(1),
		ValidArgs: common.ScenarioCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := common.Client().FindScenario(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			play, err := common.Client().StartPlay(cmd.Context(), &client.StartPlayRequest{
				Scenario: s.Scenario.Name,
			})
			if err != nil {
				return err
			}

			fmt.Println(play)

			return nil
		},
	}

	common.ScenarioFlags(cmd)

	return cmd
}
