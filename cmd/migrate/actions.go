package main

import (
	"github.com/smantic/sqlt/migrate"
	"github.com/urfave/cli/v2"
)

func CreateAction(c *cli.Context) error {
	var (
		dsn = getDSN(c)
	)
	_ = dsn
	return nil
}
func UpAction(c *cli.Context) error {

	var (
		dsn = getDSN(c)
		ctx = c.Context
	)

	return migrate.Up(ctx, dsn, nil)
}
func DownAction(c *cli.Context) error {
	var (
		dsn = getDSN(c)
		ctx = c.Context
	)
	_ = dsn

	return migrate.Down(ctx)
}
func DropAction(c *cli.Context) error {
	var (
		dsn = getDSN(c)
	)
	_ = dsn

	return nil
}

func getDSN(c *cli.Context) string {

	var dsn = c.String("dsn")

	if dsn == "" {
		dsn = c.Args().First()
	}

	return dsn
}
