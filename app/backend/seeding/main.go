package main

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
	"log"

	"gorm.io/gorm"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

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
					log.Printf("Failed to seed blood type %v: %v", bt, err)
				} else {
					log.Printf("Seeded blood type: %v", bt)
				}
			} else {
				log.Printf("DB error: %v", err)
			}
		}
	}
}

func main() {
	SeedBloodTypes()
}
