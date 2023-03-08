package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shoriwe/rollee-test-2/common/postgres"
	"github.com/shoriwe/rollee-test-2/common/sqlite"
	"github.com/shoriwe/rollee-test-2/controller"
	"github.com/shoriwe/rollee-test-2/handler"
	"github.com/urfave/cli/v2"
)

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func fromFlags(ctx *cli.Context) *controller.Controller { // Database
	dbUrl := ctx.String("database")
	switch {
	case strings.Index(dbUrl, "sqlite://") == 0:
		return controller.New(sqlite.New(dbUrl[len("sqlite://"):]))
	case strings.Index(dbUrl, "postgres://") == 0:
		return controller.New(postgres.New(dbUrl[len("postgres://"):]))
	default:
		exitErr(fmt.Errorf("expecting sqlite://filename?args or \"postgres://dsn\" for database URL"))
	}
	panic("imposible path")
}

func main() {
	app := &cli.App{
		Name:  "pluto",
		Usage: "Market data caching",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"db"},
				Value:   "sqlite://file:in_memory?mode=memory&cache=shared",
				Usage:   "Database to use",
			},
		},
		Action: func(ctx *cli.Context) error {
			c := fromFlags(ctx)
			h := handler.New(c)
			if ctx.Args().Len() > 0 {
				return h.Run(ctx.Args().First())
			}
			return h.Run("127.0.0.1:8000")
		},
	}
	app.Run(os.Args)
}
