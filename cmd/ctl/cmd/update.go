package cmd

import (
	"fmt"

	"github.com/happsie/fivem-loader/internal"
	"github.com/happsie/fivem-loader/internal/config"
	"github.com/happsie/fivem-loader/internal/updater"
	"github.com/thoas/go-funk"
	"github.com/urfave/cli/v2"
)

func Update() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		scriptName := ctx.String("script-name")
		conf, err := config.LoadConfig()
		if err != nil {
			return err
		}
		if scriptName != "" {
			is := funk.Find(conf.InstalledScripts, func(is config.InstalledScript) bool {
				return is.Name == scriptName
			})
			if is == nil {
				return fmt.Errorf("script with name %s not found", scriptName)
			}
			return update(conf.ServerDataPath, is.(config.InstalledScript))
		}
		encountedError := false
		for _, is := range conf.InstalledScripts {
			err := update(conf.ServerDataPath, is)
			if err != nil {
				encountedError = true
				fmt.Printf(internal.ErrorColor, fmt.Sprintf("could not update script [%s]: %v\n", is.Name, err))
			}
		}
		if encountedError {
			fmt.Printf(internal.WarningColor, "Scripts updated. Some script(s) produced an error, please check console output\n")
		} else {
			fmt.Printf(internal.InfoColor, "All scripts updated\n")
		}
		return nil
	}
}

func update(serverDataPath string, is config.InstalledScript) error {
	scriptUpdater := updater.ScriptUpdater{
		ServerDataPath: serverDataPath,
	}
	err := scriptUpdater.Update(is.Name, is.Github, is.ResourceFolder, is.SkippedConfig, true)
	if err != nil {
		return err
	}
	fmt.Printf(internal.InfoColor, fmt.Sprintf("script updated [%s]\n", is.Name))
	return nil
}
