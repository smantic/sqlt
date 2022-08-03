package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {

	ctx := context.Background()

	ctx, _ = signal.NotifyContext(ctx)

	app := cli.App{
		Name:                 "migrate",
		HelpName:             "",
		Usage:                "database migration tool",
		UsageText:            "",
		ArgsUsage:            "",
		Version:              "0.1",
		Description:          "",
		DefaultCommand:       "",
		Commands:             commands(),
		Flags:                nil,
		EnableBashCompletion: false,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          true,
		//Action:                 func(*cli.Context) error { panic("not implemented") },

		Compiled:               time.Now(),
		Authors:                []*cli.Author{{"tyler beverley", "tyler@smantic.dev"}},
		Copyright:              "",
		Metadata:               nil,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
		Suggest:                false,
	}

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
