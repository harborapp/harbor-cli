package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/umschlag/umschlag-cli/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	if env := os.Getenv("UMSCHLAG_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:     "umschlag-cli",
		Version:  version.String,
		Usage:    "docker distribution management system",
		Compiled: time.Now(),

		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "server, s",
				Value:   "http://localhost:8080",
				Usage:   "api server",
				EnvVars: []string{"UMSCHLAG_SERVER"},
			},
			&cli.StringFlag{
				Name:    "token, t",
				Value:   "",
				Usage:   "api token",
				EnvVars: []string{"UMSCHLAG_TOKEN"},
			},
		},

		Commands: []*cli.Command{
			Profile(),
			Registry(),
			Tag(),
			Repo(),
			Org(),
			User(),
			Team(),
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
