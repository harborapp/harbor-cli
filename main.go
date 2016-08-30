package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/umschlag/umschlag-cli/cmd"
	"github.com/umschlag/umschlag-cli/config"
	"github.com/urfave/cli"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "umschlag-cli"
	app.Version = config.Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A docker distribution management system"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server, s",
			Value:  "",
			Usage:  "Umschlag API server",
			EnvVar: "UMSCHLAG_SERVER",
		},
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			Usage:  "Umschlag API token",
			EnvVar: "UMSCHLAG_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		cmd.Profile(),
		cmd.Registry(),
		cmd.Tag(),
		cmd.Repo(),
		cmd.Org(),
		cmd.User(),
		cmd.Team(),
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
