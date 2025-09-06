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
}

func (c *Config) Load() {
	c.TextbeltAPIKey = os.Getenv("API_KEY")
	c.TextbeltURL = os.Getenv("TEXTBELT_URL")
	c.PrimalBloodlinePhoneNumber = os.Getenv("PHONE_NUMBER")
	// TODO: add logic to detect whether port is present or not
	c.Port, _ = strconv.Atoi(os.Getenv("PORT"))
}
