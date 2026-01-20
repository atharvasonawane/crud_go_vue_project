package db

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Report struct {
		Header struct {
			Institution string `yaml:"institution"`
			Address     string `yaml:"address"`
			Contact     string `yaml:"contact"`
		} `yaml:"header"`
		Footer struct {
			GeneratedBy string `yaml:"generated_by"`
		} `yaml:"footer"`
	} `yaml:"report"`
}

var AppConfig Config

func LoadConfig() {
	cwd, _ := os.Getwd()
	configPath := filepath.Join(cwd, "config", "config.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		// Keep this error! It helps if the file is missing later.
		fmt.Printf("[WARNING] Could not read config.yaml: %v\n", err)
		return
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		fmt.Printf("[ERROR] Could not parse config.yaml: %v\n", err)
		return
	}

	// Success! (Silently return, or print just one line if you prefer)
	fmt.Println("Configuration loaded successfully")
}