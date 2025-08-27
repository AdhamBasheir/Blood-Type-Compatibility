package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	BloodTypeID uint      `gorm:"not null"` // Foreign key
	BloodType   BloodType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
