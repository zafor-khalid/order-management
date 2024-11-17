package migrations

import (
	"order-management/internal/models"

	"gorm.io/gorm"
)

// Migrate handles database schema migrations for creating, modifying, or updating tables
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Order{},
	)
	if err != nil {
		return err
	}
	return nil
}
