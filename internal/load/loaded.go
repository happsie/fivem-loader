package load

import (
	"fmt"

	"github.com/happsie/fivem-loader/internal/setup"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
)

func Loaded() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		conf, err := setup.LoadConfig()
		if err != nil {
			return err
		}
		if len(conf.InstalledScripts) == 0 {
			fmt.Println("No loaded scripts found")
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
