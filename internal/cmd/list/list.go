/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package list

import (
	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/metadata"
	"github.com/srepio/cli/internal/views/list"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List the available practice scenarios",
		RunE: func(cmd *cobra.Command, args []string) error {
			md, err := metadata.Get()
			if err != nil {
				return err
			}

			tbl := list.NewTable(md)
			tbl.Print()

			return nil
		},
	}
}
