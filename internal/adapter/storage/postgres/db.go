package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/adapter/config"
)

// New creates a new PostgreSQL database instance
func New(ctx context.Context, config *config.DB) (*gorm.DB, error) {
	switch config.Connection {
	case "postgres":
		return NewPostgres(config)
	default:
		return NewPostgres(config)
	}
}
