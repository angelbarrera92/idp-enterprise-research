package main

import (
	"angelbarrera92/dagger-chainguard/cli"
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	if err := cli.Root(&ctx, os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
