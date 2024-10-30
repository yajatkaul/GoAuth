package main

import (
	"GoAuth/backend/controllers"
	"GoAuth/backend/inits"

	"github.com/gin-gonic/gin"
)

//Special function
func init() {
	inits.LoadEnvVaribles()
	inits.SetupDatabase()
}

func main() {
	r := gin.Default()
	//Auth routes
	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.Login)
	/*
	r.GET("/api/logout", controllers.LogOut)
	*/
	r.Run() // listen and serve on 0.0.0.0:PORT
}