package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	DB         PostgresConfig
	ServerPort string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DB: PostgresConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_EXTERNAL_PORT"),
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
		},
	}

	return cfg, nil
}
