package main

import (
	"GoAuth/backend/inits"
	"GoAuth/backend/models"
)

func init() {
	inits.LoadEnvVaribles()
	inits.ConnectToDb()
}

func main() {
    err := inits.DB.AutoMigrate(&models.User{})
    if err != nil {
        panic("failed to migrate database schema")
    }
}