package inits

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB_URI")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
}