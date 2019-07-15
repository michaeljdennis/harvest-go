package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config stores the Harvest account ID and token that
// are loaded to the environment from the .env file.
type Config struct {
	HarvestAccountID string
	HarvestToken     string
}

// New loads the env vars from the .env file and returns
// a Config pointer with the Harvest account ID and token.
func New(envPath string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		HarvestAccountID: os.Getenv("HARVEST_ACCOUNT_ID"),
		HarvestToken:     os.Getenv("HARVEST_TOKEN"),
	}, nil
}
