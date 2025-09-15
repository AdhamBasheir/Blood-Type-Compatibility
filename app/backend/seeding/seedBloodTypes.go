package main

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SeedBloodTypes() {
	bloodTypes := []models.BloodType{
		{ABO: "A", Rh: true},
		{ABO: "A", Rh: false},
		{ABO: "B", Rh: true},
		{ABO: "B", Rh: false},
		{ABO: "AB", Rh: true},
		{ABO: "AB", Rh: false},
		{ABO: "O", Rh: true},
		{ABO: "O", Rh: false},
	}

	for _, bt := range bloodTypes {
		var existing models.BloodType
		// Check if it exists to avoid duplicates on repeated runs
		err := initializers.DB.Where("abo = ? AND rh = ?", bt.ABO, bt.Rh).First(&existing).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := initializers.DB.Create(&bt).Error; err != nil {
					logrus.WithError(err).Errorf("Failed to seed blood type %v", bt)
					logrus.WithFields(logrus.Fields{
						"abo": bt.ABO,
						"rh":  bt.Rh,
					}).Error("Failed to seed blood type")
				}
			} else {
				logrus.WithError(err).Error("Error checking existing blood type")
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"abo": bt.ABO,
				"rh":  bt.Rh,
			}).Info("Blood type already exists, skipping seeding")
		}
	}
}
