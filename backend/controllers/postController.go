package controllers

import (
	"GoChat/backend/inits"
	"GoChat/backend/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignUpBody struct {
    UserName string `json:"userName"`
    Password string `json:"password"`
    ConfirmPassword string `json:"confirmPassword"`
}

func SignUp(c *gin.Context) {
    // Parse the raw JSON body
    var body SignUpBody
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if (body.Password != body.ConfirmPassword){
        c.JSON(http.StatusBadRequest, gin.H{"error": "Password's dont match"})
        return
    }

    user := models.User{
        UserName: body.UserName,
        Password: body.Password,
    }

    result := inits.DB.Create(&user)
    if result.Error != nil {
        if errors.Is(result.Error,  gorm.ErrDuplicatedKey) {
            // Handle unique constraint violation
            c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
        } else {
            // Handle other errors
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        }
        return
    }

    // Set a cookie
    cookieName := "session_id" 
    cookieValue := "example_session_value"
    httpOnly := true
    secure := false
    
    c.SetCookie(cookieName, cookieValue, 3600, "/", "", secure, httpOnly)

    c.JSON(http.StatusOK, gin.H{"user": user})
}

type LogInBody struct {
    UserName string `json:"userName"`
    Password string `json:"password"`
}

func LogIn(c *gin.Context) {
    // Parse the raw JSON body
    var body LogInBody
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    result := inits.DB.Where("user_name = ?", body.UserName).First(&user)

    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User or password incorrect"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        }
        return
    }

    // Set a cookie
    cookieName := "session_id" 
    cookieValue := "example_session_value"
    httpOnly := true
    secure := false
    
    c.SetCookie(cookieName, cookieValue, 3600, "/", "", secure, httpOnly)

    c.JSON(http.StatusOK, gin.H{"user": user.UserName})
}

func LogOut(c *gin.Context) {   
    cookieName := "session_id"

    // Set the cookie with an expired date to remove it
    c.SetCookie(cookieName, "", -1, "/", "", false, true)

    // Respond to the client
    c.JSON(http.StatusOK, gin.H{"message": "Cookie removed"})
}

