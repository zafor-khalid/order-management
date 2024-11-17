package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all the configuration values
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	Env        string
	JWTSecret  string
}

// AppConfig is the global variable to access configuration across the app
var AppConfig Config

// LoadConfig loads the configuration from the environment variables or .env file
func LoadConfig() error {
	// Load environment variables from .env file if present
	err := godotenv.Load()
	if err != nil {
		if getEnv("ENV", "development") == "development" {
			log.Printf("Warning: .env file not found: %v", err)
		} else {
			return fmt.Errorf("error loading .env file: %w", err)
		}
	}

	// Populate AppConfig struct with environment variables
	AppConfig = Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "order_db"),
		ServerPort: getEnv("PORT", "8080"),
		Env:        getEnv("ENV", "development"),
		JWTSecret:  getEnv("JWT_SECRET", ""),
	}
	
	return nil
}

// getEnv gets environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
