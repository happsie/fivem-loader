package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/happsie/fivem-loader/internal"
	"github.com/happsie/fivem-loader/internal/config"
	"github.com/happsie/fivem-loader/internal/updater"
	"github.com/urfave/cli/v2"
)

func Load() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		scriptName := ctx.String("script-name")
		url := getUrl(ctx.String("url"))
		if url == "" {
			return fmt.Errorf("invalid url")
		}
		config, err := config.LoadConfig()
		if err != nil {
			return err
		}
		scriptUpdater := updater.ScriptUpdater{
			ServerDataPath: config.ServerDataPath,
		}
		resourceFolders, err := getResourceFolders(config.ServerDataPath)
		if err != nil {
			return err
		}
		selectedFolder, err := Select("Select resource folder", resourceFolders)
		if err != nil {
			return err
		}
		err = scriptUpdater.Update(scriptName, url, selectedFolder, ctx.Bool("skip-cfg"), false)
		if err != nil {
			return err
		}
		fmt.Printf(internal.InfoColor, "Installation of script complete\n")
		return nil
	}
}

func getResourceFolders(path string) ([]string, error) {
	directories, err := os.ReadDir(filepath.Join(path, "resources"))
	if err != nil {
		return nil, err
	}
	var folderNames []string
	for _, dir := range directories {
		if dir.IsDir() {
			folderNames = append(folderNames, dir.Name())
		}
	}
	return folderNames, nil
}

func getUrl(url string) string {
	if strings.Contains(url, "github.com") {
		return fmt.Sprintf("%s/archive/refs/heads/master.zip", url)
	}
	return url
}
