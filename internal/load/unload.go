package load

import (
	"fmt"
	"os"

	"github.com/happsie/fivem-loader/internal/setup"
	"github.com/thoas/go-funk"
	"github.com/urfave/cli/v2"
)

func Unload() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		scriptName := ctx.String("script-name")
		if scriptName == "" {
			return fmt.Errorf("specify script name")
		}
		conf, err := setup.LoadConfig()
		if err != nil {
			return err
		}
		script := funk.Find(conf.InstalledScripts, func(is setup.InstalledScript) bool {
			return is.Name == scriptName
		})
		if script == nil {
			return fmt.Errorf("script not found")
		}
		err = os.RemoveAll(script.(setup.InstalledScript).Location)
		if err != nil {
			return err
		}
		conf.InstalledScripts = funk.Filter(conf.InstalledScripts, func(is setup.InstalledScript) bool {
			return is.Name != scriptName
		}).([]setup.InstalledScript)
		err = conf.Save()
		if err != nil {
			return err
		}
		fmt.Printf("Script unloaded [%s]", scriptName)
		return nil
	}
}
