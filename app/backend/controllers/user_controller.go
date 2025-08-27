package controllers

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ListUsers(ctx *gin.Context) {
	// Placeholder for user retrieval logic
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
		log.Printf("Failed to create user %v: %v", newUser, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	} else {
		log.Printf("Created user: %v", newUser)
		ctx.JSON(http.StatusOK, gin.H{"user": newUser})
	}
}
