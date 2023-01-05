package cmd

import (
	"fmt"
	"os"

	"github.com/happsie/fivem-loader/internal"
	"github.com/happsie/fivem-loader/internal/config"
	"github.com/thoas/go-funk"
	"github.com/urfave/cli/v2"
)

func Unload() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		scriptName := ctx.String("script-name")
		if scriptName == "" {
			return fmt.Errorf("specify script name")
		}
		conf, err := config.LoadConfig()
		if err != nil {
			return err
		}
		script := funk.Find(conf.InstalledScripts, func(is config.InstalledScript) bool {
			return is.Name == scriptName
		})
		if script == nil {
			return fmt.Errorf("script not found")
		}
		err = os.RemoveAll(script.(config.InstalledScript).Location)
		if err != nil {
			return err
		}
		conf.InstalledScripts = funk.Filter(conf.InstalledScripts, func(is config.InstalledScript) bool {
			return is.Name != scriptName
		}).([]config.InstalledScript)
		err = conf.Save()
		if err != nil {
			return err
		}
		fmt.Printf(internal.InfoColor, fmt.Sprintf("Script unloaded [%s]", scriptName))
		return nil
	}
}
