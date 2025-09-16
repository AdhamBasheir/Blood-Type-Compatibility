package services

import (
	"blood-type-compatibility/helpers"
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
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
		Name:        input.Name,
		Username:    input.Username,
		Password:    hashed,
		BloodTypeID: input.BloodTypeID,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
