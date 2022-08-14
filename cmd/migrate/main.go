package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/urfave/cli/v2"
)

// Version is the most recent tag for the repo. Injected by the build
var Version string

func main() {

	ctx, _ := signal.NotifyContext(context.Background())

	app := cli.App{
		Name:                   "migrate",
		Usage:                  "database migration tool",
		UsageText:              "",
		ArgsUsage:              "",
		Version:                Version,
		Description:            "",
		DefaultCommand:         "",
		Commands:               commands,
		Flags:                  nil,
		EnableBashCompletion:   false,
		HideHelp:               false,
		HideHelpCommand:        false,
		HideVersion:            true,
		Authors:                []*cli.Author{{"tyler beverley", "tyler@smantic.dev"}},
		Copyright:              "",
		Metadata:               nil,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
		Suggest:                false,
	}

	if err := app.RunContext(ctx, os.Args); err != nil {
		fmt.Println(err)
	}
}
