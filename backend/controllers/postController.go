package controllers

import (
	"GoChat/backend/inits"
	"GoChat/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpBody struct {
    UserName string `json:"userName"`
    Password string `json:"password"`
}

func SignUp(c *gin.Context) {
    // Parse the raw JSON body
    var body SignUpBody
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{
        UserName: body.UserName,
        Password: body.Password,
    }

    result := inits.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}
