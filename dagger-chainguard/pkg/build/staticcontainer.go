package build

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func BuildStaticContainer(ctx *context.Context, name string, binaries ...string) error {
	// initialize Dagger client
	client, err := dagger.Connect(*ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	pwd := client.Host().Directory(".")
	container := client.Container().From("cgr.dev/chainguard/static:latest")
	for _, binary := range binaries {
		container = container.WithFile(binary, pwd.File(binary))
	}

	img, err := container.Publish(*ctx, name, dagger.ContainerPublishOpts{})
	if err != nil {
		return err
	}
	fmt.Println("Published image", img)
	return nil
}
