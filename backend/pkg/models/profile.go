package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID         uint         `gorm:"column:user_id"`
	Name           string       `gorm:"column:name"`
	Bio            string       `gorm:"column:bio"`
	ProfilePicture string       `gorm:"column:profile_picture"`
	BirthDate      *time.Time   `gorm:"column:birth_date"`    // Menggunakan pointer untuk mengizinkan null
	Projects       []Project    `gorm:"foreignKey:ProfileID"` // Foreign key to Profile
	Skills         []Skill      `gorm:"foreignKey:ProfileID"` // Foreign key to Profile
	Contacts       []Contact    `gorm:"foreignKey:ProfileID"` // Foreign key to Profile
	Educations     []Education  `gorm:"foreignKey:ProfileID"` // Foreign key to Profile
	Experiences    []Experience `gorm:"foreignKey:ProfileID"` // Foreign key to Profile
}
