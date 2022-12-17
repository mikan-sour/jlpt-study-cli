package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_Host     string
	DB_Port     string
	DB_Name     string
	DB_Username string
	DB_Password string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Name:     os.Getenv("DB_NAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Username: os.Getenv("DB_USERNAME"),
	}
}
