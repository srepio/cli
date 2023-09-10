/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package kill

import (
	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
)

func NewKillCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "kill [scenario]",
		Short:     "Kill the specified practice scenario",
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
			return d.Kill(cmd.Context(), instance)
		},
	}

	common.ScenarioFlags(cmd)

	return cmd
}
