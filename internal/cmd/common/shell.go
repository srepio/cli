package common

import (
	"context"
	"os"

	"github.com/srepio/sdk/client"
	"golang.org/x/term"
)

func RunShell(play string) error {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	ctx := context.Background()
	req := &client.GetShellRequest{
		ID: play,
	}
	return Client().GetShell(ctx, req, os.Stdin, os.Stdout)
}
