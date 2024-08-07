package inits

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVaribles() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}