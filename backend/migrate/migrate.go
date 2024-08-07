package main

import (
	"GoChat/backend/inits"
	"GoChat/backend/models"
)

func init() {
	inits.LoadEnvVaribles()
	inits.ConnectToDb()
}

func main() {
	inits.DB.AutoMigrate(&models.User{})
}