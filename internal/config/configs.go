package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Configs struct {
	Database DbConfig
}

func NewConfigs() (Configs, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Configs{}, err
	}

	return Configs{
		Database: DbConfig{
			DbConnectionString: os.Getenv(os.Getenv("DB_DSN")),
		},
	}, nil
}
