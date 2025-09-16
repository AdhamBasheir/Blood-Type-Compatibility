package controllers

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
	"blood-type-compatibility/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ListUsers(ctx *gin.Context) {
	var users []models.User
	if err := initializers.DB.Find(&users).Error; err != nil {
		logrus.WithError(err).Error("Failed to fetch users")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(ctx *gin.Context) {
	// Placeholder for user retrieval logic
}

func CreateUser(ctx *gin.Context) {
	// Placeholder for user creation logic
}

func UpdateUser(ctx *gin.Context) {
	// Placeholder for updating user information
}

func DeleteUser(ctx *gin.Context) {
	// Placeholder for deleting a user
}

func SignUp(ctx *gin.Context) {
	var body struct {
		Name        string `json:"name" binding:"required"`
		Username    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		BloodTypeID uint   `json:"blood_type_id" binding:"required"`
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request"})
		return
	}

	user, err := services.CreateUser(services.UserPayLoad{
		Name:        body.Name,
		Username:    body.Username,
		Password:    body.Password,
		BloodTypeID: body.BloodTypeID,
	})
	if err != nil {
		logrus.WithError(err).Error("Failed to create user")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"username": user.Username,
		"user_id":  user.ID,
	}).Info("User created successfully")

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
