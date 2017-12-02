package main

import (
	"os"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/umschlag/umschlag-cli/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("UMSCHLAG_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:     "umschlag-cli",
		Version:  version.Version.String(),
		Usage:    "Docker distribution management system",
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
				Usage:   "Umschlag API server",
				EnvVars: []string{"UMSCHLAG_SERVER"},
			},
			&cli.StringFlag{
				Name:    "token, t",
				Value:   "",
				Usage:   "Umschlag API token",
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
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
