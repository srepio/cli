/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package check

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/sdk/client"
	"github.com/srepio/sdk/types"
)

func NewCheckCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "check [scenario]",
		Short:     "Check the specified practice scenario",
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

			if d.Check(cmd.Context(), instance) {
				fmt.Println("The check script passed!")
				play, err := d.Kill(cmd.Context(), instance)
				if err != nil {
					return err
				}
				if _, err := common.Client().CompletePlay(cmd.Context(), &client.CompletePlayRequest{ID: play}); err != nil {
					return err
				}
			}

			fmt.Println("The check script failed, try again")

			return nil
		},
	}

	common.ScenarioFlags(cmd)

	return cmd
}
