package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv from .env file
func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	return os.Getenv(key)
}
