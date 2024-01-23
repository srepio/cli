/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package ctx

import (
	"github.com/cqroot/prompt"
	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
)

func NewCtxCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ctx",
		Short:   "Update the connection that SREP CLI uses",
		GroupID: "auth",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctxs := []string{}
			for _, conn := range common.Config.Connections {
				ctxs = append(ctxs, conn.Name)
			}
			ctx, err := prompt.New().Choose(ctxs)
			if err != nil {
				return err
			}

			return common.Config.SetContext(ctx)
		},
	}
	return cmd
}
