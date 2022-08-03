package main

import "github.com/urfave/cli/v2"

// statically declared commands.
func commands() []*cli.Command {
	var commands []*cli.Command = []*cli.Command{
		&cli.Command{
			Name:        "create",
			Aliases:     []string{},
			Usage:       "Create a new migration",
			UsageText:   "",
			Description: "",
			ArgsUsage:   "",
			Category:    "",

			SkipFlagParsing:        false,
			HideHelp:               false,
			HideHelpCommand:        false,
			Hidden:                 false,
			UseShortOptionHandling: false,
			HelpName:               "",
			CustomHelpTemplate:     "",

			Subcommands: []*cli.Command{},
			Flags:       []cli.Flag{},
			Action:      CreateAction,
		},
		&cli.Command{
			Name:        "up",
			Aliases:     []string{},
			Usage:       "Apply all up migrations ",
			UsageText:   "migrate up $DSN [--source] [--file]  ",
			Description: "",
			ArgsUsage:   "",
			Category:    "",

			SkipFlagParsing:        false,
			HideHelp:               false,
			HideHelpCommand:        false,
			Hidden:                 false,
			UseShortOptionHandling: false,
			HelpName:               "",
			CustomHelpTemplate:     "",

			Subcommands: []*cli.Command{},
			Flags:       []cli.Flag{dsnFlag, sourceFlag, fileFlag},
			Action:      UpAction,
		},
		&cli.Command{
			Name:                   "down",
			Aliases:                []string{},
			Usage:                  "Apply all down migrations",
			UsageText:              "migrate down $DSN [--source] [--file]",
			Description:            "",
			ArgsUsage:              "",
			Category:               "",
			SkipFlagParsing:        false,
			HideHelp:               false,
			HideHelpCommand:        false,
			Hidden:                 false,
			UseShortOptionHandling: false,
			HelpName:               "",
			CustomHelpTemplate:     "",

			Subcommands: []*cli.Command{},
			Flags:       []cli.Flag{dsnFlag, sourceFlag, fileFlag},
			Action:      DownAction,
		},
		&cli.Command{
			Name:                   "drop",
			Aliases:                []string{},
			Usage:                  "Drop all schemas in the database",
			UsageText:              "migrate drop $DSN",
			Description:            "",
			ArgsUsage:              "DSN = data source name of your database.",
			Category:               "",
			SkipFlagParsing:        false,
			HideHelp:               false,
			HideHelpCommand:        false,
			Hidden:                 false,
			UseShortOptionHandling: false,
			HelpName:               "",
			CustomHelpTemplate:     "",

			Subcommands: []*cli.Command{},
			Flags:       []cli.Flag{dsnFlag, sourceFlag, fileFlag},
			Action:      DropAction,
		},
	}

	return commands
}
