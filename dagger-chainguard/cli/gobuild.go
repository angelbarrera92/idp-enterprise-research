package cli

import (
	"angelbarrera92/dagger-chainguard/pkg/build"
	"angelbarrera92/dagger-chainguard/pkg/build/gobuilder"
	"context"
	"flag"
	"fmt"
	"os"
)

type GoBuildCommand struct {
	ctx *context.Context

	fs *flag.FlagSet

	output string
	main   string
	os     string
	arch   string

	oci_name string
}

func (g *GoBuildCommand) Name() string {
	return g.fs.Name()
}

func (g *GoBuildCommand) Init(ctx *context.Context, args []string) error {
	g.ctx = ctx
	return g.fs.Parse(args)
}

func (g *GoBuildCommand) Run() error {
	// Create GoBuilderOptions struct
	opts := &gobuilder.GoBuilderOptions{
		Main: g.main,
		OS:   g.os,
		Arch: g.arch,
		Out:  g.output,
	}
	// Create GoBuilder struct
	builder := gobuilder.NewGoBuilder(g.ctx)
	err := builder.Build(opts)
	if err != nil {
		return err
	}
	build.BuildStaticContainer(g.ctx, g.oci_name, g.output)

	return nil
}

func (g *GoBuildCommand) Validate() error {
	// Check main file exists, if not, return error
	if _, err := os.Stat(g.main); os.IsNotExist(err) {
		return fmt.Errorf("main file %s does not exist", g.main)
	}
	// Check output file doesn't exist, if it does, return error
	if _, err := os.Stat(g.output); !os.IsNotExist(err) {
		return fmt.Errorf("output file %s already exists", g.output)
	}
	// Check if os is a golang valid os, if not, return error
	// Currently, only linux binaries are supported as we use linux containers
	if g.os != "linux" {
		return fmt.Errorf("os %s is not supported", g.os)
	}
	// Check if arch is a valid golang arch, if not, return error
	// Currently, only amd64 or arch64 binaries are supported as we use linux containers
	if g.arch != "amd64" && g.arch != "arm64" {
		return fmt.Errorf("arch %s is not supported", g.arch)
	}
	// Check if oci_name is a valid oci name, if not, return error
	if g.oci_name == "" {
		return fmt.Errorf("oci_name %s is not a valid oci name", g.oci_name)
	}

	return nil
}

func NewGoBuildCommand() *GoBuildCommand {
	gc := &GoBuildCommand{
		fs: flag.NewFlagSet("go-build", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.output, "o", "app", "output file")
	gc.fs.StringVar(&gc.main, "main", "main.go", "main package")
	gc.fs.StringVar(&gc.os, "os", "linux", "target operating system")
	gc.fs.StringVar(&gc.arch, "arch", "amd64", "target architecture")
	gc.fs.StringVar(&gc.oci_name, "oci-name", "", "name of the container image to be published")

	return gc
}
