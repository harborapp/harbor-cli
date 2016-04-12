package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/harborapp/harbor-cli/cmd"
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"github.com/harborapp/harbor-cli/config"
)

var (
	updates string = "http://dl.webhippie.de/"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "harbor-cli"
	app.Version = config.Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A docker distribution management system"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server, s",
			Value:  "",
			Usage:  "Solder API server",
			EnvVar: "HARBOR_SERVER",
		},
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			Usage:  "Solder API token",
			EnvVar: "HARBOR_TOKEN",
		},
		cli.BoolFlag{
			Name:   "update, u",
			Usage:  "Enable auto update",
			EnvVar: "HARBOR_UPDATE",
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.Bool("update") {
			if config.VersionDev == "dev" {
				fmt.Fprintf(os.Stderr, "Updates are disabled for development versions.\n")
			} else {
				updater := &selfupdate.Updater{
					CurrentVersion: config.StrippedVersion,
					ApiURL:         updates,
					BinURL:         updates,
					DiffURL:        updates,
					Dir:            "updates/",
					CmdName:        app.Name,
				}

				go updater.BackgroundRun()
			}
		}

		return nil
	}

	app.Commands = []cli.Command{
		cmd.Profile(),
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
