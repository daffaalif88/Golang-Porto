package models

import (
	"gorm.io/gorm"
)

// SetupDatabase initializes the database connection and migrates the schema
func SetupDatabase(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(
		&User{},
		&Profile{},
		&Project{},
		&Skill{},
		&Contact{},
		&Education{},
		&Experience{},
	)
}
