package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type TenantGetter interface {
	Tenant() port.TenantRepository
}

type TenantRepository struct {
	db *gorm.DB
}

func newTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{
		db: db,
	}
}

func (r *TenantRepository) Create(ctx context.Context, data *domain.Tenant) (*domain.Tenant, error) {
	if err := r.db.Model(&domain.Tenant{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *TenantRepository) List(ctx context.Context, req *domain.TenantListRequest) ([]*domain.Tenant, int, error) {
	var datas []*domain.Tenant
	err := r.db.Model(&domain.Tenant{}).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, 0, nil
}

func (r *TenantRepository) Get(ctx context.Context, id string) (*domain.Tenant, error) {
	var data domain.Tenant
	if err := r.db.Model(&domain.Tenant{}).
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *TenantRepository) Update(ctx context.Context, id string, req domain.Map) (*domain.Tenant, error) {
	data := &domain.Tenant{}
	err := r.db.Model(&domain.Tenant{}).Where("id = ?", id).Updates(req).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *TenantRepository) Delete(ctx context.Context, id string) error {
	return r.db.Model(&domain.Tenant{}).Where("id = ?", id).Delete(&domain.Tenant{}).Error
}
