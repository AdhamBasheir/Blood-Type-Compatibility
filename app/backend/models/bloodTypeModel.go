package models

import "gorm.io/gorm"

type BloodType struct {
	gorm.Model
	ABO string `gorm:"not null"` // A, B, AB, O
	Rh  bool   `gorm:"not null"` // true for positive, false for negative
	// User []User
}
