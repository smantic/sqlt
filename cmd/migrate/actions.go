package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/smantic/sqlt/migrate"
	"github.com/urfave/cli/v2"
)

var (
	MissingDriverScheme = errors.New(
		"expected driver scheme to be presnt in the DSN. Use --driver or specify the driver in the DSN (driver://dsn)",
	)
	MissingDSN = errors.New("specify a DSN to connect to with --dsn")
)

func validateDSN(c *cli.Context) error {

	var dsn = c.String("dsn")

	if dsn == "" {
		dsn = c.Args().First()
	}

	if dsn == "" {
		return MissingDSN
	}

	if driver := c.String("driver"); driver != "" {
		index := strings.Index(dsn, "://")
		if index == -1 {
			dsn = fmt.Sprintf("%s://%s", driver, dsn)
		} else {
			// otherwise replace the existing driver scheme.
			dsn = fmt.Sprintf("%s://%s", driver, dsn[index+2:])
		}
	}

	if !strings.Contains(dsn, "://") {
		return MissingDriverScheme
	}

	return c.Set("dsn", dsn)
}

func CreateAction(c *cli.Context) error {
	var (
		dsn = c.String("dsn")
	)
	_ = dsn
	return nil
}

func UpAction(c *cli.Context) error {

	var (
		dsn = c.String("dsn")
		ctx = c.Context
	)

	return migrate.Up(ctx, dsn, nil)
}

func DownAction(c *cli.Context) error {
	var (
		dsn = c.String("dsn")
		ctx = c.Context
	)

	return migrate.Down(ctx, dsn)
}
func DropAction(c *cli.Context) error {
	var (
		dsn = c.String("dsn")
	)
	_ = dsn

	return nil
}
