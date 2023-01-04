package main

import (
	"log"
	"os"

	"github.com/happsie/fivem-loader/internal/load"
	"github.com/happsie/fivem-loader/internal/setup"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "setup",
				Usage: "First time setup for FiveM-Loader. Required as a first step before installing scripts",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "server-data-path",
						Aliases:  []string{"sdp"},
						Required: true,
					},
				},
				Action: setup.Setup(),
			},
			{
				Name:  "load",
				Usage: "Loads script to selected resource and ensures script in server.cfg",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "github",
						Aliases:  []string{"g"},
						Usage:    "Link to github repository containing script. (example: https://github.com/happsie/fivem-hello-world)",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "script-name",
						Aliases:  []string{"name", "sn"},
						Usage:    "Name of the script that will also be used for creating the resource folder. (example: hello_world)",
						Required: true,
					},
					&cli.BoolFlag{
						Name:     "skip-cfg",
						Aliases:  []string{"scfg"},
						Usage:    "Skips addition to server.cfg",
						Required: false,
						Value:    false,
					},
				},
				Action: load.Load(),
			},
			{
				Name:   "loaded",
				Usage:  "List loaded scripts by FiveM Loader",
				Action: load.Loaded(),
			},
			{
				Name:  "unload",
				Usage: "unloads selected script",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "script-name",
						Aliases:  []string{"name", "sn"},
						Usage:    "Script name to uninstall. (example: hello_world)",
						Required: true,
					},
				},
				Action: load.Unload(),
			},
		},
	}
	app.Name = "FiveM-Loader"
	app.Authors = []*cli.Author{
		{
			Name: "happsie",
		},
	}
	app.Usage = "Fast and easy installation of scripts"
	app.Description = "A cli application for installing FiveM scripts"
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
