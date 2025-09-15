package main

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedTestUsers() {
	// Example names just for testing
	baseNames := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}

	// There are 8 blood types
	totalBloodTypes := 8
	password := "password"

	for i := 1; i <= 64; i++ {
		name := fmt.Sprintf("%s%d", baseNames[i%len(baseNames)], i)
		username := fmt.Sprintf("user%d", i)

		// Pick blood type in round-robin fashion
		bloodTypeID := uint((i % totalBloodTypes) + 1)

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			logrus.WithError(err).Error("Failed to hash password for test user")
			continue
		}

		user := models.User{
			Name:        name,
			Username:    username,
			Password:    string(hash),
			BloodTypeID: bloodTypeID,
		}

		// Check if user already exists to avoid duplicates on repeated runs
		var existing models.User
		err = initializers.DB.Where("username = ?", username).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := initializers.DB.Create(&user).Error; err != nil {
				logrus.WithFields(logrus.Fields{
					"username":      username,
					"blood_type_id": bloodTypeID,
				}).WithError(err).Error("Failed to seed test user")
			} else {
				logrus.WithFields(logrus.Fields{
					"username":      username,
					"blood_type_id": bloodTypeID,
				}).Info("Seeded test user")
			}
		} else if err != nil {
			logrus.WithFields(logrus.Fields{
				"username":      username,
				"blood_type_id": bloodTypeID,
			}).WithError(err).Error("Failed to check if user exists")
		} else {
			logrus.WithFields(logrus.Fields{
				"username":      username,
				"blood_type_id": bloodTypeID,
			}).Info("User already exists, skipping seed")
		}
	}
}
