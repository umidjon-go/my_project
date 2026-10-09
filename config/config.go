package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresDb   string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASS"),
		PostgresDb:   os.Getenv("POSTGRES_DB"),
	}
}
