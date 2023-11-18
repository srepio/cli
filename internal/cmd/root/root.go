/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package root

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/check"
	"github.com/srepio/cli/internal/cmd/describe"
	"github.com/srepio/cli/internal/cmd/initialise"
	"github.com/srepio/cli/internal/cmd/kill"
	"github.com/srepio/cli/internal/cmd/list"
	"github.com/srepio/cli/internal/cmd/run"
	"github.com/srepio/cli/internal/config"
)

var (
	Config *config.Config
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
			Config = c
			return nil
		},
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	cmd.SetVersionTemplate(fmt.Sprintf("%s version %s commit %s built at %s\n", cmd.Use, version, commit, date))

	cmd.AddCommand(list.NewListCommand())
	cmd.AddCommand(run.NewRunCommand())
	cmd.AddCommand(check.NewCheckCommand())
	cmd.AddCommand(describe.NewDescribeCommand())
	cmd.AddCommand(kill.NewKillCommand())
	cmd.AddCommand(initialise.NewInitCommand())

	return cmd
}
