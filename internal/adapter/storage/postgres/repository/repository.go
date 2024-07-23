package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/core/port"
)

type Repository struct {
	db *gorm.DB
}

type IRepository interface {
	UserGetter
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) User() port.UserRepository {
	return newUserRepository(r.db)
}
