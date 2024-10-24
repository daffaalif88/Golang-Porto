package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ProfileID       uint   `gorm:"not null"`
	Title           string `gorm:"type:varchar(100)"`
	Description     string `gorm:"type:text"`
	TechnologyStack string `gorm:"type:varchar(255)"`
	Image           string `gorm:"type:varchar(255)"`
	ProjectURL      string `gorm:"type:varchar(255)"`
	DateCreated     string `gorm:"type:date"`
}
