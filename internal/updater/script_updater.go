package updater

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/happsie/fivem-loader/internal/setup"
	"github.com/thoas/go-funk"
)

const comment = "# FiveM-Loader scripts"

type ScriptUpdater struct {
	ServerDataPath string
}

// Update takes scriptName (name of the script), url (github url) and destination folder (the resource folder to place the script in e.g [local])
func (su *ScriptUpdater) Update(scriptName, url, destinationFolder string, skipConfig bool) error {
	if su.ServerDataPath == "" {
		return fmt.Errorf("server data path not setup, try setting it up before loading scripts")
	}
	conf, err := setup.LoadConfig()
	if err != nil {
		return err
	}
	alreadyInstalled := funk.Contains(conf.InstalledScripts, func(is setup.InstalledScript) bool {
		return is.Name == scriptName && strings.Contains(is.Location, destinationFolder)
	})
	if alreadyInstalled {
		return fmt.Errorf("script already installed with name %s", scriptName)
	}
	zipName, err := DownloadZip(getGithubLink(url))
	if err != nil {
		return err
	}
	defer RemoveFile(zipName)
	err = Unzip(su.getResourcesPath(destinationFolder), scriptName, zipName)
	if err != nil {
		return err
	}
	if skipConfig == false {
		su.updateServerCfg(scriptName)
	}
	conf.InstalledScripts = append(conf.InstalledScripts, setup.InstalledScript{
		Name:     scriptName,
		Github:   url,
		Location: filepath.Join(su.getResourcesPath(destinationFolder), scriptName),
	})
	err = conf.Save()
	if err != nil {
		return err
	}
	return nil
}

func (su ScriptUpdater) updateServerCfg(scriptName string) error {
	b, err := os.ReadFile(su.getCfgPath())
	if err != nil {
		return err
	}
	serverCfg := string(b)
	if !strings.Contains(serverCfg, comment) {
		serverCfg += fmt.Sprintf("\n\n%s", comment)
	}
	if !strings.Contains(serverCfg, fmt.Sprintf("ensure %s", scriptName)) {
		serverCfg += fmt.Sprintf("\nensure %s", scriptName)
	}
	err = os.WriteFile(su.getCfgPath(), []byte(serverCfg), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (su ScriptUpdater) getResourcesPath(destinationFolder string) string {
	return fmt.Sprintf(`%s\resources\%s\`, su.ServerDataPath, destinationFolder)
}

func (su ScriptUpdater) getCfgPath() string {
	if strings.HasSuffix(su.ServerDataPath, "\\") {
		return fmt.Sprintf(`%sserver.cfg`, su.ServerDataPath)
	}
	return fmt.Sprintf(`%s\server.cfg`, su.ServerDataPath)
}

func getGithubLink(url string) string {
	return fmt.Sprintf("%s/archive/refs/heads/master.zip", url)
}
