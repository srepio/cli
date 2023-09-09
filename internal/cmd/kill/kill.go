/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package kill

import (
	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/cli/internal/metadata"
)

var (
	k8s   bool
	clean bool
)

func NewKillCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "kill [scenario]",
		Short:     "Kill the specified practice scenario",
		Args:      cobra.ExactArgs(1),
		ValidArgs: common.ScenarioCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := metadata.Find(args[0])
			if err != nil {
				return err
			}

			d, err := common.GetDriver(k8s)
			if err != nil {
				return err
			}
			instance, err := d.Create(*s)
			if err != nil {
				return err
			}
			return d.Kill(cmd.Context(), instance)
		},
	}

	cmd.Flags().BoolVar(&k8s, "k8s", false, "Determine whether to use kubernetes as the driver. Defaults to using docker.")
	cmd.Flags().BoolVar(&clean, "clean", true, "Determine whether to kill and remove the instance.")

	return cmd
}
