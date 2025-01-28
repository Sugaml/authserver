package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type ClientSecretGetter interface {
	ClientSecret() port.ClientSecretRepository
}

type ClientSecretRepository struct {
	db *gorm.DB
}

func newClientSecretRepository(db *gorm.DB) *ClientSecretRepository {
	return &ClientSecretRepository{
		db: db,
	}
}

func (r *ClientSecretRepository) Create(ctx context.Context, data *domain.ClientSecret) (*domain.ClientSecret, error) {
	if err := r.db.Model(&domain.ClientSecret{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ClientSecretRepository) List(ctx context.Context, req *domain.ClientSecretListRequest) ([]*domain.ClientSecret, int, error) {
	var datas []*domain.ClientSecret
	err := r.db.Model(&domain.ClientSecret{}).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, 0, nil
}

func (r *ClientSecretRepository) ListByApplicationID(ctx context.Context, id string, req *domain.ClientSecretListRequest) ([]*domain.ClientSecret, int, error) {
	var datas []*domain.ClientSecret
	count := 0
	f := r.db.Model(&domain.ClientSecret{})
	if req.Query != "" {
		f = f.Where("lower(title) LIKE lower(?) and lower(subtitle) LIKE lower(?)", "%"+req.Query+"%", "%"+req.Query+"%")
	}
	err := f.
		Where("application_id = ?", id).
		Count(&count).
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, nil
}

func (r *ClientSecretRepository) ListByClientID(ctx context.Context, id string) ([]*domain.ClientSecret, int, error) {
	var datas []*domain.ClientSecret
	err := r.db.Model(&domain.ClientSecret{}).Where("client_id = ?", id).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, 0, nil
}

func (r *ClientSecretRepository) Get(ctx context.Context, id string) (*domain.ClientSecret, error) {
	var data domain.ClientSecret
	if err := r.db.Model(&domain.ClientSecret{}).
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ClientSecretRepository) GetClientIDAndValue(ctx context.Context, clientID, value string) (*domain.ClientSecret, error) {
	var data domain.ClientSecret
	err := r.db.Model(&domain.ClientSecret{}).Take(&data, "client_id = ? AND value", clientID, value).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ClientSecretRepository) Update(ctx context.Context, id string, req domain.Map) (*domain.ClientSecret, error) {
	data := &domain.ClientSecret{}
	err := r.db.Model(&domain.ClientSecret{}).Where("id = ?", id).Updates(req).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ClientSecretRepository) UpdateIsActive(ctx context.Context, id string, isActive bool) (*domain.ClientSecret, error) {
	data := &domain.ClientSecret{}
	err := r.db.Model(&domain.ClientSecret{}).
		Where("id = ?", id).
		UpdateColumns(
			map[string]interface{}{
				"is_active": isActive,
			},
		).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ClientSecretRepository) Delete(ctx context.Context, id string) error {
	return r.db.Model(&domain.ClientSecret{}).Where("id = ?", id).Delete(&domain.ClientSecret{}).Error
}
