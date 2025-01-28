package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type ApplicationGetter interface {
	Application() port.ApplicationRepository
}

type ApplicationRepository struct {
	db *gorm.DB
}

func newApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}

func (r *ApplicationRepository) Create(ctx context.Context, data *domain.Application) (*domain.Application, error) {
	if err := r.db.Model(&domain.Application{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ApplicationRepository) List(ctx context.Context, req *domain.ListApplicationRequest) ([]*domain.Application, int, error) {
	var datas []*domain.Application
	err := r.db.Model(&domain.Application{}).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, 0, nil
}

func (r *ApplicationRepository) Get(ctx context.Context, id string) (*domain.Application, error) {
	var data domain.Application
	if err := r.db.Model(&domain.Application{}).
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ApplicationRepository) Update(ctx context.Context, id string, req domain.Map) (*domain.Application, error) {
	data := &domain.Application{}
	err := r.db.Model(&domain.Application{}).Where("id = ?", id).Updates(req).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ApplicationRepository) Delete(ctx context.Context, id string) error {
	logrus.Info("package repository Delete() Application function called.")
	return r.db.Model(&domain.Application{}).Where("id = ?", id).Delete(&domain.Application{}).Error
}
