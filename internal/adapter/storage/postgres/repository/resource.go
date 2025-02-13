package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type ResourceGetter interface {
	Resource() port.ResourceRepository
}

type ResourceRepository struct {
	db *gorm.DB
}

func newResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{
		db: db,
	}
}

func (r *ResourceRepository) Create(ctx context.Context, data *domain.Resource) (*domain.Resource, error) {
	if err := r.db.Model(&domain.Resource{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ResourceRepository) List(ctx context.Context, req *domain.ResourceListRequest) ([]*domain.Resource, int, error) {
	var datas []*domain.Resource
	err := r.db.Model(&domain.Resource{}).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, 0, nil
}

func (r *ResourceRepository) Get(ctx context.Context, id string) (*domain.Resource, error) {
	var data domain.Resource
	if err := r.db.Model(&domain.Resource{}).
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ResourceRepository) Update(ctx context.Context, id string, req domain.Map) (*domain.Resource, error) {
	data := &domain.Resource{}
	err := r.db.Model(&domain.Resource{}).Where("id = ?", id).Updates(req).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ResourceRepository) Delete(ctx context.Context, id string) error {
	return r.db.Model(&domain.Resource{}).Where("id = ?", id).Delete(&domain.Resource{}).Error
}
