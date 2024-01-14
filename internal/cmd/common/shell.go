package common

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/srepio/sdk/client"
	"golang.org/x/term"
)

func RunShell(play string) error {
	oldState, tErr := term.MakeRaw(int(os.Stdin.Fd()))
	if tErr != nil {
		return tErr
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	ctx := context.Background()
	req := &client.GetShellRequest{
		ID: play,
	}

	var err error
	tries := 0
	for err == nil {
		if tries > 3 {
			return errors.New("unable to connect to play after 3 attempts")
		}
		err = Client().GetShell(ctx, req, os.Stdin, os.Stdout)
		if errors.Is(err, client.ErrTooEarly) {
			err = nil
			tries++
			time.Sleep(time.Second * 3)
		} else {
			return err
		}
	}
	return errors.New("fallthrough shell error")
}
