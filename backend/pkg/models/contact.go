package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ProfileID    uint   `gorm:"not null"`
	ContactType  string `gorm:"type:varchar(50)"`
	ContactValue string `gorm:"type:varchar(255)"`
}
