/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package root

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/auth/ctx"
	"github.com/srepio/cli/internal/cmd/auth/login"
	"github.com/srepio/cli/internal/cmd/cancel"
	"github.com/srepio/cli/internal/cmd/check"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/cli/internal/cmd/describe"
	"github.com/srepio/cli/internal/cmd/initialise"
	"github.com/srepio/cli/internal/cmd/list"
	"github.com/srepio/cli/internal/cmd/run"
	"github.com/srepio/cli/internal/cmd/shell"
	"github.com/srepio/cli/internal/config"
)

func BuildRootCmd(version, commit, date string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "srep",
		Short:   "SRE practice",
		Long:    `A CLI that runs SRE practice scenarios`,
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			c, err := config.GetConfig()
			if err != nil {
				if errors.Is(err, config.ErrNoConfig) {
					fmt.Printf("Config file not found.\nRun `srep init` to generate a new config\n")
					os.Exit(0)
				}
				return err
			}
			common.Config = c
			common.InitClient(c)
			return nil
		},
	}

	cmd.SetVersionTemplate(fmt.Sprintf("%s version %s commit %s built at %s\n", cmd.Use, version, commit, date))

	cmd.AddGroup(&cobra.Group{ID: "play", Title: "Play commands:"})
	cmd.AddGroup(&cobra.Group{ID: "auth", Title: "Auth commands:"})
	cmd.AddGroup(&cobra.Group{ID: "other", Title: "Other commands:"})

	cmd.SetHelpCommandGroupID("other")
	cmd.SetCompletionCommandGroupID("other")

	cmd.AddCommand(list.NewListCommand())
	cmd.AddCommand(run.NewRunCommand())
	cmd.AddCommand(check.NewCheckCommand())
	cmd.AddCommand(describe.NewDescribeCommand())
	cmd.AddCommand(cancel.NewCancelCommand())
	cmd.AddCommand(initialise.NewInitCommand())
	cmd.AddCommand(shell.NewShellCommand())

	cmd.AddCommand(login.NewLoginCommand())
	cmd.AddCommand(ctx.NewCtxCommand())

	return cmd
}
