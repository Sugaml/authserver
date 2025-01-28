package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/core/domain"
)

// Migrate up database table
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
		&domain.Customer{},
		&domain.Client{},
		&domain.ClientSecret{},
	).Error
}
