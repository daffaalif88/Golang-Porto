package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string  `gorm:"type:varchar(255);not null"`
	Password string  `gorm:"type:varchar(255);not null"`
	Profile  Profile `gorm:"foreignKey:UserID"` // One-to-One relationship
}
