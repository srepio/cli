/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package cancel

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/sdk/client"
)

func NewCancelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cancel",
		Short:   "Cancel the active play",
		GroupID: "srep",
		RunE: func(cmd *cobra.Command, args []string) error {
			active, err := common.Client().GetActivePlay(cmd.Context(), &client.GetActivePlayRequest{})
			if err != nil {
				return err
			}

			if active.Play == nil {
				fmt.Println("No active play")
				return nil
			}

			_, err = common.Client().CancelPlay(cmd.Context(), &client.CancelPlayRequest{
				ID: active.Play.ID,
			})

			return err
		},
	}

	common.ScenarioFlags(cmd)

	return cmd
}
