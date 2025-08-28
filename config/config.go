package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	TextbeltAPIKey string `json:"textbelt_api_key"`
}

func (c *Config) Load(filePath string) error {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileBytes, c)
	if err != nil {
		return err
	}
	return nil
}
