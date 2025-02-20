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
	CustomerGetter
	ClientGetter
	ClientSecretGetter
	ApplicationGetter
	RoleGetter
	ResourceGetter
	TenantGetter
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) User() port.UserRepository {
	return newUserRepository(r.db)
}

func (r *Repository) Customer() port.CustomerRepository {
	return newCustomerRepository(r.db)
}

func (r *Repository) Client() port.ClientRepository {
	return newClientRepository(r.db)
}

func (r *Repository) ClientSecret() port.ClientSecretRepository {
	return newClientSecretRepository(r.db)
}

func (r *Repository) Application() port.ApplicationRepository {
	return newApplicationRepository(r.db)
}

func (r *Repository) Tenant() port.TenantRepository {
	return newTenantRepository(r.db)
}

func (r *Repository) Resource() port.ResourceRepository {
	return newResourceRepository(r.db)
}

func (r *Repository) Role() port.RoleRepository {
	return newRoleRepository(r.db)
}
