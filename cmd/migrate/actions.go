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
	_ = dsn

	return migrate.Up(ctx, nil)
}
func DownAction(c *cli.Context) error {
	var (
		dsn = getDSN(c)
		ctx = c.Context
	)
	_ = dsn
	return migrate.Down(ctx, nil)
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
		args := c.Args().First()
	}

	return dsn
}
