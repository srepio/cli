/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package shell

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/sdk/client"
)

func NewShellCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "shell [scenario]",
		Short:   "Get a shell in the active play",
		GroupID: "play",
		RunE: func(cmd *cobra.Command, args []string) error {
			active, err := common.Client().GetActivePlay(cmd.Context(), &client.GetActivePlayRequest{})
			if err != nil {
				return err
			}

			if active.Play == nil {
				return errors.New("no active play running")
			}

			return common.RunShell(cmd.Context(), active.Play.ID)
		},
	}

	common.ScenarioFlags(cmd)

	return cmd
}
