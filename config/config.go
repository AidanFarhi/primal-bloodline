package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	TextbeltAPIKey             string `json:"textbelt_api_key"`
	TextbeltURL                string `json:"textbelt_url"`
	PrimalBloodlinePhoneNumber string `json:"primal_bloodline_phone_number"`
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
