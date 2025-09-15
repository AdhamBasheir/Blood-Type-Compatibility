package controllers

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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
		UserName    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		BloodTypeID uint   `json:"blood_type_id" binding:"required"`
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := models.User{
		Name:        body.Name,
		UserName:    body.UserName,
		Password:    string(hash),
		BloodTypeID: body.BloodTypeID,
	}

	if err := initializers.DB.Create(&newUser).Error; err != nil {
		logrus.WithError(err).Error("Failed to create user")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	} else {
		logrus.WithFields(logrus.Fields{
			"username": newUser.UserName,
			"user_id":  newUser.ID,
		}).Info("User created successfully")
		ctx.JSON(http.StatusOK, gin.H{"user": newUser})
	}
}
