package internal

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

var (
	LASTFM_KEY    = ""
	LASTFM_SECRET = ""
)

type Config struct {
	Session struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	}
}

func createConfigFile(filepath string) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()
}

func LoadConfig() (Config, bool) {
	config_dir, _ := os.UserConfigDir()
	app_dir := path.Join(config_dir, "Scrobbleme")

	configFilePath := path.Join(app_dir, "config.json")
	var config Config

	file, err := os.Open(configFilePath)
	if err != nil {
		createConfigFile(configFilePath)
		return config, false
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, false
	}

	return config, true
}

func SaveConfig(config Config) {
	config_dir, _ := os.UserConfigDir()
	app_dir := path.Join(config_dir, "Scrobbleme")

	configFilePath := path.Join(app_dir, "config.json")

	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(configFilePath, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
