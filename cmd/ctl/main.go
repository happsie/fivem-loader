package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/happsie/fivem-loader/internal/setup"
	"github.com/happsie/fivem-loader/internal/updater"
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
				Usage: "Loads script to [local] and ensures script in server.cfg",
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
						Usage:    "Name of the folder (script name) that will be created inside [local]. (example: hello_world)",
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					scriptName := ctx.String("script-name")
					if scriptName == "" {
						return fmt.Errorf("script name cannot be empty")
					}
					github := ctx.String("github")
					if github == "" && !strings.HasPrefix(github, "https://github.com") {
						return fmt.Errorf("invalid github url")
					}
					config, err := setup.LoadConfig()
					if err != nil {
						return err
					}
					scriptUpdater := updater.ScriptUpdater{
						ServerDataPath: config.ServerDataPath,
					}
					err = scriptUpdater.Update(scriptName, github)
					if err != nil {
						return err
					}
					log.Println("Script installation complete")
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
