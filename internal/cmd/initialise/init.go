/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package initialise

import (
	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/config"
)

func NewInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "init",
		Short:   "Initialise the config",
		GroupID: "other",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return config.Initialise()
		},
	}
}
