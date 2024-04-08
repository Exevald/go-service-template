package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"
)

const appName = "service-template"

func main() {
	ctx := context.Background()
	err := runApp(ctx, os.Args)
	if err != nil {
		panic(err.Error())
	}
}

func runApp(ctx context.Context, args []string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	app := cli.App{Name: appName}

	return app.RunContext(ctx, args)
}
