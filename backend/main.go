package main

import (
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:PORT
}