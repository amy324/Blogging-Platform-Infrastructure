package config

import (
	"os"
	"github.com/joho/godotenv"
)

// DBConfig holds the configuration settings for the CockroachDB database
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewDBConfig creates a new instance of DBConfig with default values
func NewDBConfig() *DBConfig {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		// Handle error if .env file is not found or cannot be read
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}

// APIConfig holds the configuration settings for the API service
type APIConfig struct {
	Port string
}

// LoadConfig loads the configuration settings for the API service
func LoadConfig() *APIConfig {
	// You can set default values here or load them from environment variables or a configuration file
	return &APIConfig{
		Port: "8080", // Default port 8080
	}
}
