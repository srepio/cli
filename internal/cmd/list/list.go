/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package list

import (
	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/cli/internal/views/list"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List the available practice scenarios",
		GroupID: "srep",
		RunE: func(cmd *cobra.Command, args []string) error {
			md, err := common.Client().Getscenarios(cmd.Context())
			if err != nil {
				return err
			}

			tbl := list.NewTable(md.Scenarios)
			tbl.Print()

			return nil
		},
	}
}
