/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package login

import (
	"context"
	"fmt"
	u "net/url"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
	"github.com/spf13/cobra"
	"github.com/srepio/cli/internal/cmd/common"
	"github.com/srepio/cli/internal/config"
	"github.com/srepio/sdk/client"
)

var (
	url            string
	connectionName string
)

func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Login to SREP",
		GroupID: "auth",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Update the client's current config to point to our new url
			u, err := u.Parse(url)
			if err != nil {
				return err
			}
			common.Client().Options.Url = u.Host
			common.Client().Options.Scheme = u.Scheme

			email, err := prompt.New().Ask("Email:").Input("")
			if err != nil {
				return err
			}
			password, err := prompt.New().Ask("Password:").Input("", input.WithEchoMode(input.EchoPassword))
			if err != nil {
				return err
			}

			resp, err := login(cmd.Context(), email, password)
			if err != nil {
				return err
			}
			// Use the JWT we got back from login to now create an api key
			common.Client().Options.Token = resp.Token

			host, err := os.Hostname()
			if err != nil {
				return err
			}
			code, err := common.Client().CreateApiToken(cmd.Context(), &client.CreateApiTokenRequest{
				Name: fmt.Sprintf("SREP CLI (%s)", host),
			})
			if err != nil {
				return err
			}

			return common.Config.SetConnection(config.Api{
				Name:   connectionName,
				Url:    u.Host,
				Scheme: u.Scheme,
				Token:  code.Token,
			})
		},
	}

	cmd.Flags().StringVarP(&url, "url", "u", "https://api.srep.io", "The url of the SREP API instance")
	cmd.Flags().StringVarP(&connectionName, "connection", "c", "default", "The name of the connection to save your credentials to")

	return cmd
}

func login(ctx context.Context, email, password string) (*client.LoginResponse, error) {
	resp, err := common.Client().Login(ctx, &client.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	if !resp.MFARequired {
		return resp, nil
	}

	code, err := prompt.New().Ask("MFA Code:").Input("", input.WithEchoMode(input.EchoPassword))
	if err != nil {
		return nil, err
	}

	return common.Client().VerifyMFA(ctx, &client.VerifyMFARequest{
		AuthenticationID: resp.AuthenticationID,
		Code:             code,
	})
}
