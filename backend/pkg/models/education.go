package models

import "gorm.io/gorm"

type Education struct {
	gorm.Model
	ProfileID    uint   `gorm:"not null"`
	Institution  string `gorm:"type:varchar(100)"`
	Degree       string `gorm:"type:varchar(100)"`
	FieldOfStudy string `gorm:"type:varchar(100)"`
	StartDate    string `gorm:"type:date"`
	EndDate      string `gorm:"type:date"`
}
