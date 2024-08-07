package main

import (
	"GoChat/backend/controllers"
	"GoChat/backend/inits"

	"github.com/gin-gonic/gin"
)

//Special function
func init() {
	inits.LoadEnvVaribles()
	inits.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.LogIn)
	r.GET("/api/logout", controllers.LogOut)
	r.Run() // listen and serve on 0.0.0.0:PORT
}