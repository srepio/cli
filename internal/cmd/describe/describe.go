/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package describe

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
)

func NewDescribeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "describe [scenario]",
		Short:     "Describe the specified practice scenario",
		Args:      cobra.ExactArgs(1),
		ValidArgs: common.ScenarioCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := common.Client().FindScenario(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			fmt.Printf("Scenario: %s\n", s.Scenario.Name)
			fmt.Printf("Difficulty: %s\n", s.Scenario.Difficulty)
			fmt.Println("Tags:")
			for _, tag := range s.Scenario.Tags {
				fmt.Printf("  - %s\n", tag)
			}
			fmt.Printf("\nDescription:\n")
			fmt.Println(s.Scenario.Description)

			return nil
		},
	}

	return cmd
}
