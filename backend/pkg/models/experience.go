package models

import "gorm.io/gorm"

type Experience struct {
	gorm.Model
	ProfileID   uint   `gorm:"not null"`
	CompanyName string `gorm:"type:varchar(100)"`
	Position    string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
	StartDate   string `gorm:"type:date"`
	EndDate     string `gorm:"type:date"`
}
