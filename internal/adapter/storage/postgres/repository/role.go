package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type RoleGetter interface {
	Role() port.RoleRepository
}

type RoleRepository struct {
	db *gorm.DB
}

func newRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) Create(ctx context.Context, data *domain.Role) (*domain.Role, error) {
	if err := r.db.Model(&domain.Role{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *RoleRepository) List(ctx context.Context, req *domain.RoleListRequest) ([]*domain.Role, int, error) {
	var datas []*domain.Role
	err := r.db.Model(&domain.Role{}).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, 0, nil
}

func (r *RoleRepository) Get(ctx context.Context, id string) (*domain.Role, error) {
	var data domain.Role
	if err := r.db.Model(&domain.Role{}).
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RoleRepository) Update(ctx context.Context, id string, req domain.Map) (*domain.Role, error) {
	data := &domain.Role{}
	err := r.db.Model(&domain.Role{}).Where("id = ?", id).Updates(req).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *RoleRepository) Delete(ctx context.Context, id string) error {
	return r.db.Model(&domain.Role{}).Where("id = ?", id).Delete(&domain.Role{}).Error
}
