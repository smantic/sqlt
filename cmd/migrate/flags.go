package main

import (
	"time"

	"github.com/urfave/cli/v2"
)

var (
	dsnFlag = &cli.StringFlag{
		Name:        "dsn",
		Category:    "",
		DefaultText: "",
		FilePath:    "",
		Usage:       "Run migrations against this data source name (driver://url)",
		Required:    false,
		Hidden:      false,
		HasBeenSet:  false,
		Value:       "",
		Destination: nil,
		Aliases:     []string{"database"},
		EnvVars:     []string{"DSN"},
		TakesFile:   false,
	}
	driverFlag = &cli.StringFlag{
		Name:        "driver",
		Category:    "",
		DefaultText: "",
		FilePath:    "",
		Usage:       "specify the driver to use with the DSN (postgres, mysql ...)",
		Required:    false,
		Hidden:      false,
		HasBeenSet:  false,
		Value:       "",
	}
	sourceFlag = &cli.StringFlag{
		Name:        "source",
		Category:    "",
		DefaultText: "",
		FilePath:    "",
		Usage:       "Path to a directory containing migrations",
		Required:    false,
		Hidden:      false,
		HasBeenSet:  false,
		Value:       "",
		Destination: nil,
		Aliases:     []string{"path"},
		EnvVars:     []string{},
		TakesFile:   false,
	}
	fileFlag = &cli.StringFlag{
		Name:        "file",
		Category:    "",
		DefaultText: "",
		FilePath:    "",
		Usage:       "Run the migrations in the specified file",
		Required:    false,
		Hidden:      false,
		HasBeenSet:  false,
		Value:       "",
		Destination: nil,
		Aliases:     []string{},
		EnvVars:     []string{},
		TakesFile:   true,
	}
	timeoutFlag = &cli.DurationFlag{
		Name:  "timeout",
		Usage: "provide a duration that the migration should run for before timeing out and rolling back.",
		Value: time.Minute * 5,
	}
)
