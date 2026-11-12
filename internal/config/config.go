package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL string
}

func NewConfig() Config {
	return Config{}
}

func (c *Config) Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	c.DB_URL = os.Getenv("DB_URL")
}
