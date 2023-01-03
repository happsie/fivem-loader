package updater

import (
	"fmt"
	"os"
	"strings"
)

const comment = "# FiveM-Loader scripts"

type ScriptUpdater struct {
	ServerDataPath string
}

func (su *ScriptUpdater) Update(scriptName, url string) error {
	if su.ServerDataPath == "" {
		return fmt.Errorf("server data path not setup, try setting it up before loading scripts")
	}
	zipName, err := DownloadZip(getGithubLink(url))
	if err != nil {
		return err
	}
	defer RemoveFile(zipName)
	err = Unzip(su.getResourcesPath(), scriptName, zipName)
	if err != nil {
		return err
	}
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

func (su ScriptUpdater) getResourcesPath() string {
	return fmt.Sprintf(`%s\resources\[local]\`, su.ServerDataPath)
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
