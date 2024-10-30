package controllers

import (
	"context"
	"time"

	"net/http"

	"GoAuth/backend/inits"
	"GoAuth/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := inits.DB.Database("testdb").Collection("users")

	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sessionId := uuid.New().String()

	expTime := 7 * 24 * time.Hour

	cookie := &http.Cookie{
		Name:    "sessionId",
		Value:   sessionId,
		Expires: time.Now().Add(expTime), 
		Path:    "/",
	}
	http.SetCookie(c.Writer, cookie)

	validTill := time.Now().Add(expTime)
	session := models.Sessions{
		ID:        sessionId, 
		UserId:    user.ID.Hex(),
		ValidTill: validTill,
	}

	sessionCollection := inits.DB.Database("testdb").Collection("sessions")
	_, err = sessionCollection.InsertOne(context.TODO(), session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var login LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	collection := inits.DB.Database("testdb").Collection("users")

	err := collection.FindOne(context.TODO(), bson.M{"username": login.Username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	if login.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	sessionId := uuid.New().String()

	expTime := 7 * 24 * time.Hour

	cookie := &http.Cookie{
		Name:    "sessionId",
		Value:   sessionId,
		Expires: time.Now().Add(expTime), 
		Path:    "/",
	}
	http.SetCookie(c.Writer, cookie)

	validTill := time.Now().Add(expTime)
	session := models.Sessions{
		ID:        sessionId, 
		UserId:    user.ID.Hex(),
		ValidTill: validTill,
	}

	sessionCollection := inits.DB.Database("testdb").Collection("sessions")
	_, err = sessionCollection.InsertOne(context.TODO(), session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})	
}

func Logout(c *gin.Context) {   
    cookieName := "sessionId"

    // Set the cookie with an expired date to remove it
    c.SetCookie(cookieName, "", -1, "/", "", false, true)

    // Respond to the client
    c.JSON(http.StatusOK, gin.H{"message": "Cookie removed"})
}