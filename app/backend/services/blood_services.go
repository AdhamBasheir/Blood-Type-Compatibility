package services

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
)

func FindCompatibleDonors(abo string, rh bool) ([]models.User, error) {
	var compatibleBloodTypes []string
	var users []models.User
	var err error

	switch abo {
	case "O":
		compatibleBloodTypes = []string{"O"}
	case "A":
		compatibleBloodTypes = []string{"O", "A"}
	case "B":
		compatibleBloodTypes = []string{"O", "B"}
	case "AB":
		compatibleBloodTypes = []string{"O", "A", "B", "AB"}
	}

	if rh {
		err = initializers.DB.InnerJoins("BloodType").Where("abo IN ?", compatibleBloodTypes).Find(&users).Error
	} else {
		err = initializers.DB.InnerJoins("BloodType").Where("abo IN ? AND rh = ?", compatibleBloodTypes, false).Find(&users).Error
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}
