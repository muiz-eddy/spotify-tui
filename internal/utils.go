package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// GetEnv this will load .env relative to the file that run this function
func GetEnv(key string) string {

	absPath, _ := filepath.Abs("../../.env")

	err := godotenv.Load(absPath)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
