package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	ProfileID   uint   `gorm:"not null"`
	SkillName   string `gorm:"type:varchar(100)"`
	Proficiency string `gorm:"type:varchar(50)"`
}
