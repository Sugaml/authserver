package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type ClientGetter interface {
	Client() port.ClientRepository
}

type ClientRepository struct {
	db *gorm.DB
}

func newClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

// NewClientRepository creates a new Client repository instance
func NewClientRepository(db *gorm.DB) port.ClientRepository {
	return &ClientRepository{
		db,
	}
}

func (r *ClientRepository) Create(ctx context.Context, data *domain.Client) (*domain.Client, error) {
	if err := r.db.Model(&domain.Client{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ClientRepository) List(ctx context.Context, req *domain.ClientListRequest) ([]*domain.Client, int, error) {
	var datas []*domain.Client
	err := r.db.Model(&domain.Client{}).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, 0, nil
}

func (r *ClientRepository) ListByApplicationID(ctx context.Context, id string, req *domain.ClientListRequest) ([]*domain.Client, int, error) {
	var datas []*domain.Client
	count := 0
	f := r.db.Model(&domain.Client{})
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

func (r *ClientRepository) Get(ctx context.Context, id string) (*domain.Client, error) {
	var data domain.Client
	if err := r.db.Model(&domain.Client{}).Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ClientRepository) GetCliendID(ctx context.Context, clientID string) (*domain.Client, error) {
	var data domain.Client
	if err := r.db.Model(&domain.Client{}).Take(&data, "client_id = ?", clientID).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ClientRepository) Update(ctx context.Context, id string, req domain.Map) (*domain.Client, error) {
	data := &domain.Client{}
	err := r.db.Model(&domain.Client{}).Where("id = ?", id).Updates(req).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ClientRepository) UpdateIsActive(ctx context.Context, id string, isActive bool) (*domain.Client, error) {
	data := &domain.Client{}
	err := r.db.Model(&domain.Client{}).
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

func (r *ClientRepository) Delete(ctx context.Context, id string) error {
	return r.db.Model(&domain.Client{}).Where("id = ?", id).Delete(&domain.Client{}).Error
}
