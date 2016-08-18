package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sanbornm/go-selfupdate/selfupdate"
	"github.com/umschlag/umschlag-cli/cmd"
	"github.com/umschlag/umschlag-cli/config"
	"github.com/urfave/cli"
)

var (
	updates = "http://dl.webhippie.de/"
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
		cli.BoolTFlag{
			Name:   "update, u",
			Usage:  "Enable auto update",
			EnvVar: "UMSCHLAG_UPDATE",
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.BoolT("update") {
			if config.VersionDev == "dev" {
				fmt.Fprintf(os.Stderr, "Updates are disabled for development versions.\n")
			} else {
				updater := &selfupdate.Updater{
					CurrentVersion: fmt.Sprintf(
						"%d.%d.%d",
						config.VersionMajor,
						config.VersionMinor,
						config.VersionPatch,
					),
					ApiURL:  updates,
					BinURL:  updates,
					DiffURL: updates,
					Dir:     "updates/",
					CmdName: app.Name,
				}

				go updater.BackgroundRun()
			}
		}

		return nil
	}

	app.Commands = []cli.Command{
		cmd.Profile(),
		cmd.Registry(),
		cmd.User(),
		cmd.Team(),
		cmd.Org(),
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app.Run(os.Args)
}
