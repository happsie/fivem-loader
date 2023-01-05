package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type InstalledScript struct {
	Name           string `json:"name"`
	Github         string `json:"github"`
	Location       string `json:"location"`
	ResourceFolder string `json:"resourceFolder"`
	SkippedConfig  bool   `json:"skippedConfig"`
}

type Config struct {
	ServerDataPath   string            `json:"serverDataPath"`
	InstalledScripts []InstalledScript `json:"installedScripts"`
}

const settingsFileName = "fivem-loader-settings.json"

func (c *Config) Save() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	settingsPath := filepath.Join(homeDir, settingsFileName)
	f, err := os.OpenFile(settingsPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	jsonByte, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(jsonByte))
	return nil
}

func CreateConfig(path string) error {
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	settingsPath := filepath.Join(homeDir, settingsFileName)
	f, err := os.Create(settingsPath)
	if err != nil {
		return err
	}
	defer f.Close()
	conf := Config{
		ServerDataPath:   path,
		InstalledScripts: []InstalledScript{},
	}
	err = conf.Save()
	if err != nil {
		return err
	}
	log.Println("config created")
	return nil
}

func LoadConfig() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	settingsPath := filepath.Join(homeDir, settingsFileName)
	jsonBytes, err := os.ReadFile(settingsPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("setup is incomplete, missing %s", settingsFileName)
	}
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
