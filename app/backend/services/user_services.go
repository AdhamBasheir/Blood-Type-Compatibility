package services

import (
	"blood-type-compatibility/helpers"
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
	"crypto/rand"
	"encoding/base64"
)

type UserPayLoad struct {
	Name        string
	Username    string
	Password    string
	BloodTypeID uint
}

func CreateUser(input UserPayLoad) (*models.User, error) {
	hashed, err := helpers.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:         input.Name,
		Username:     input.Username,
		Password:     hashed,
		BloodTypeID:  input.BloodTypeID,
		SessionToken: "",
		CSRFToken:    "",
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func AuthenticateUser(username, password string) (*models.User, bool) {
	var user models.User
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, false
	}

	if !helpers.CheckPassword(user.Password, password) {
		return nil, false
	}
	return &user, true
}

func CreateSession(userID uint) (string, string, error) {
	sessionToken, err := generateToken(helpers.SessionTokenLength)
	if err != nil {
		return "", "", err
	}

	csrfToken, err := generateToken(helpers.CSRFTokenLength)
	if err != nil {
		return "", "", err
	}

	return sessionToken, csrfToken, nil
}

func InvalidateSession(username string) error {
	var user models.User
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}

	user.SessionToken = ""
	user.CSRFToken = ""
	if err := initializers.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func generateToken(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func ValidateToken(username, sessionToken, csrfToken string) (*models.User, bool) {
	var user models.User
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, false
	}

	if user.SessionToken == sessionToken && user.CSRFToken == csrfToken && sessionToken != "" && csrfToken != "" {
		return &user, true
	}
	return nil, false
}

func AuthorizeUser(username, csrfToken, sessionToken string) (*models.User, bool) {
	user, ok := ValidateToken(username, sessionToken, csrfToken)

	if !ok {
		return nil, false
	}
	return user, true
}
