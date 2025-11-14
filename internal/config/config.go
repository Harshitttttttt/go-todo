package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        int
	DatabaseUrl string
	Environment string
	JWTSecret   string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env variables: ", err)
	}

	cfg := &Config{}

	// Port
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid Port value: %s", portStr)
	}
	cfg.Port = port

	// Database URL
	cfg.DatabaseUrl = os.Getenv("DATABASE_URL")

	// Environment
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	cfg.Environment = env

	// JWT Secret
	cfg.JWTSecret = os.Getenv("JWT_SECRET")

	return cfg
}
