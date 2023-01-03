package setup

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	ServerDataPath string `json:"serverDataPath"`
}

const settingsFileName = "fivem-loader-settings.json"

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
		ServerDataPath: path,
	}
	jsonByte, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(jsonByte))
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
