package cmd

import (
	"fmt"

	"github.com/happsie/fivem-loader/internal"
	"github.com/happsie/fivem-loader/internal/config"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
)

func Loaded() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		conf, err := config.LoadConfig()
		if err != nil {
			return err
		}
		if len(conf.InstalledScripts) == 0 {
			fmt.Printf(internal.InfoColor, "No loaded scripts found")
			return nil
		}
		tbl := table.New("Name", "Location", "Github")
		for _, script := range conf.InstalledScripts {
			tbl.AddRow(script.Name, script.Location, script.Github)
		}
		tbl.Print()
		return nil
	}
}
