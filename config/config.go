package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	HarvestAccountID string
	HarvestToken     string
}

// New ...
func New(envPath string) Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	return Config{
		HarvestAccountID: os.Getenv("HARVEST_ACCOUNT_ID"),
		HarvestToken:     os.Getenv("HARVEST_TOKEN"),
	}
}
