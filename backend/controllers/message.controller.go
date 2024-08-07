package controllers

import (
	"GoChat/backend/inits"
	"GoChat/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetChats(c *gin.Context) {
    cookieValue, err := c.Cookie("session_id")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cookie not found"})
        return
    }

    senderID, err := strconv.ParseUint(cookieValue, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
        return
    }

	var convos []models.Conversation
    result := inits.DB.Where("user_ids @> ?", "["+strconv.FormatUint(senderID, 10)+"]").Find(&convos)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"user": convos})
}