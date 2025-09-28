package controllers

import (
	"blood-type-compatibility/helpers"
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

func Login(ctx *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request"})
		return
	}

	user, ok := services.AuthenticateUser(body.Username, body.Password)
	if !ok {
		logrus.Warn("Authentication failed")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	sessionToken, csrfToken, err := services.CreateSession(user.ID)
	if err != nil {
		logrus.WithError(err).Error("Failed to create session")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	ctx.SetCookie("session_token", sessionToken, helpers.SessionDuration, "/", "localhost", false, true)
	ctx.SetCookie("csrf_token", csrfToken, helpers.SessionDuration, "/", "localhost", false, false)

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	if err := initializers.DB.Save(&user).Error; err != nil {
		logrus.WithError(err).Error("Failed to save session tokens")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session tokens"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"username": user.Username,
		"user_id":  user.ID,
	}).Info("User logged in successfully")

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"csrf_token":    csrfToken,
		"session_token": sessionToken,
	})
}

func Logout(ctx *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request"})
		return
	}

	csrfToken := ctx.GetHeader("X-CSRF-Token")
	sessionToken, err := ctx.Cookie("session_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No session token found"})
		return
	}

	user, ok := services.AuthorizeUser(body.Username, csrfToken, sessionToken)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid CSRF token or session"})
		return
	}

	ctx.SetCookie("session_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("csrf_token", "", -1, "/", "localhost", false, false)
	if err := services.InvalidateSession(user.Username); err != nil {
		logrus.WithError(err).Error("Failed to invalidate session")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"username": user.Username,
		"user_id":  user.ID,
	}).Info("User logged out successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
