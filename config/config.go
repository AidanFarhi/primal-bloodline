package config

import (
	"os"
	"strconv"
)

type Config struct {
	TextbeltAPIKey             string
	TextbeltURL                string
	PrimalBloodlinePhoneNumber string
	Port                       int
	Develop                    bool
	JSONPath                   string
}

func (c *Config) Load() {
	c.TextbeltAPIKey = os.Getenv("API_KEY")
	c.TextbeltURL = os.Getenv("TEXTBELT_URL")
	c.PrimalBloodlinePhoneNumber = os.Getenv("PHONE_NUMBER")
	c.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	c.Develop, _ = strconv.ParseBool(os.Getenv("DEVELOP"))
	c.JSONPath = os.Getenv("JSON_PATH")
}
