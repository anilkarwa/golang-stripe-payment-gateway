package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Simple helper function to read an environment
func GetEnv(key string) string {
	return os.Getenv(key)
}
