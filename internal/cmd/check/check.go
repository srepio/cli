/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package check

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
)

var (
	clean bool
)

func NewCheckCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "check [scenario]",
		Short:     "Check the specified practice scenario",
		Args:      cobra.ExactArgs(1),
		ValidArgs: common.ScenarioCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := common.Client().FindScenario(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			d, err := common.GetDriver(cmd.Flag("driver").Value.String())
			if err != nil {
				return err
			}
			instance, err := d.Create(*s.Scenario)
			if err != nil {
				return err
			}

			if d.Check(cmd.Context(), instance) {
				fmt.Println("The check script passed!")
				if clean {
					return d.Kill(cmd.Context(), instance)
				}
				return nil
			}

			fmt.Println("The check script failed, try again")

			return nil
		},
	}

	cmd.Flags().BoolVar(&clean, "clean", true, "Determine whether to kill and remove the instance.")
	common.ScenarioFlags(cmd)

	return cmd
}
