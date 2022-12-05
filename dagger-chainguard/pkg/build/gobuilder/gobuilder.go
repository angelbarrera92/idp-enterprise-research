package gobuilder

import (
	"context"
	"os"

	"dagger.io/dagger"
)

type GoBuilder struct {
	ctx *context.Context
}

type GoBuilderOptions struct {
	Main string
	OS   string
	Arch string
	Out  string
}

func NewGoBuilder(ctx *context.Context) *GoBuilder {
	return &GoBuilder{
		ctx: ctx,
	}
}

func (g *GoBuilder) Build(opts *GoBuilderOptions) error {
	// initialize Dagger client
	client, err := dagger.Connect(*g.ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	golang := client.Container().From("cgr.dev/chainguard/go:latest")

	// This is the default behavior
	work := client.Host().Directory(".")

	golang = golang.WithMountedDirectory("/work", work)
	golang = golang.WithWorkdir("/work")
	golang = golang.WithEnvVariable("GOOS", opts.OS)
	golang = golang.WithEnvVariable("GOARCH", opts.Arch)
	golang = golang.WithEnvVariable("CGO_ENABLED", "0")

	golang = golang.WithExec([]string{"build", "-o", opts.Out, opts.Main})

	output := golang.File(opts.Out)
	_, err = output.Export(*g.ctx, opts.Out)
	if err != nil {
		return err
	}

	return nil
}
