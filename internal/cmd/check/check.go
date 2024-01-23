/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package check

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/sdk/client"
)

func NewCheckCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check",
		Short:   "Check the active scenario",
		GroupID: "play",
		RunE: func(cmd *cobra.Command, args []string) error {
			active, err := common.Client().GetActivePlay(cmd.Context(), &client.GetActivePlayRequest{})
			if err != nil {
				return err
			}

			out, err := common.Client().CheckPlay(cmd.Context(), &client.CheckPlayRequest{ID: active.Play.ID})
			if err != nil {
				return err
			}

			if out.Passed {
				fmt.Println("Play passed!")
			} else {
				fmt.Println("The check script failed, try again")
			}

			return nil
		},
	}

	common.ScenarioFlags(cmd)

	return cmd
}
