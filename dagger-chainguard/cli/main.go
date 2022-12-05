package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
)

type Runner interface {
	Init(*context.Context, []string) error
	Validate() error
	Run() error
	Name() string
}

func Root(ctx *context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a sub-command")
	}

	cmds := []Runner{
		NewGoBuildCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			err := cmd.Init(ctx, os.Args[2:])
			if err != nil {
				return err
			}
			err = cmd.Validate()
			if err != nil {
				return err
			}
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
