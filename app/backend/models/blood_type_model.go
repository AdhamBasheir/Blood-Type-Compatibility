package models

import (
	"gorm.io/gorm"
)

type BloodType struct {
	gorm.Model
	ABO     string `gorm:"not null"` // A, B, AB, O
	Rh      bool   `gorm:"not null"` // true for positive, false for negative
	Display string `gorm:"-" json:"display"`
}

func (b *BloodType) AfterFind(tx *gorm.DB) (err error) {
	if b.Rh {
		b.Display = b.ABO + "+"
	} else {
		b.Display = b.ABO + "-"
	}
	return
}
