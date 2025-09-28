package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string    `gorm:"not null"`
	Username     string    `gorm:"uniqueIndex;not null"`
	Password     string    `gorm:"not null"`
	BloodTypeID  uint      `gorm:"not null"` // Foreign key
	BloodType    BloodType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SessionToken string    `gorm:"uniqueIndex"`
	CSRFToken    string    `gorm:"uniqueIndex"`
}
