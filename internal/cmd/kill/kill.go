/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package kill

import (
	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/sdk/client"
	"github.com/srepio/sdk/types"
)

func NewKillCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "kill [scenario]",
		Short:     "Kill the specified practice scenario",
		GroupID:   "srep",
		Args:      cobra.ExactArgs(1),
		ValidArgs: common.ScenarioCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := common.Client().FindScenario(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			d, err := common.GetDriver(types.DriverName(cmd.Flag("driver").Value.String()))
			if err != nil {
				return err
			}

			instance, err := d.Create(*s.Scenario)
			if err != nil {
				return err
			}
			play, err := d.Kill(cmd.Context(), instance)
			if err != nil {
				return err
			}

			if _, err := common.Client().FailPlay(cmd.Context(), &client.FailedPlayRequest{ID: play}); err != nil {
				return err
			}

			return nil
		},
	}

	common.ScenarioFlags(cmd)

	return cmd
}
