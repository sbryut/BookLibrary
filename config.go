package main

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	DB   PostgresConfig
	Port string
}

type PostgresConfig struct {
	Username string
	Password string
	Port     string
	DBName   string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: os.Getenv("SERVER_PORT"),
		DB: PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Port:     os.Getenv("POSTGRES_EXTERNAL_PORT"),
			DBName:   os.Getenv("POSTGRES_DB"),
		},
	}

	return cfg, nil
}
