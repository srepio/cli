/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package run

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/sdk/client"
	"golang.org/x/term"
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "run [scenario]",
		Short:     "Run the specified practice scenarios",
		GroupID:   "srep",
		Args:      cobra.ExactArgs(1),
		ValidArgs: common.ScenarioCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			active, err := common.Client().GetActivePlay(cmd.Context(), &client.GetActivePlayRequest{})
			if err != nil {
				return err
			}

			var playID string
			if active.Play == nil {
				id, err := startPlay(cmd.Context(), args[0])
				if err != nil {
					return err
				}
				playID = id
			} else {
				playID = active.Play.ID
			}

			oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
			if err != nil {
				return err
			}
			defer term.Restore(int(os.Stdin.Fd()), oldState)

			ctx := context.Background()
			req := &client.GetShellRequest{
				ID: playID,
			}
			if err := common.Client().GetShell(ctx, req, os.Stdin, os.Stdout); err != nil {
				return err
			}

			return nil
		},
	}

	common.ScenarioFlags(cmd)

	return cmd
}

func startPlay(ctx context.Context, name string) (string, error) {
	s, err := common.Client().FindScenario(ctx, name)
	if err != nil {
		return "", err
	}

	play, err := common.Client().StartPlay(ctx, &client.StartPlayRequest{
		Scenario: s.Scenario.Name,
	})
	if err != nil {
		return "", err
	}
	return play.ID, nil
}
