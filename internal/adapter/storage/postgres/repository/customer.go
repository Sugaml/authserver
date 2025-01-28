package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type CustomerGetter interface {
	Customer() port.CustomerRepository
}

type CustomerRepository struct {
	db *gorm.DB
}

func newCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) Create(ctx context.Context, data *domain.Customer) (*domain.Customer, error) {
	if err := r.db.Model(&domain.Customer{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CustomerRepository) List(ctx context.Context, req *domain.ListCustomerRequest) ([]*domain.Customer, int, error) {
	var datas []*domain.Customer
	err := r.db.Model(&domain.Customer{}).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, 0, nil
}

func (r *CustomerRepository) Get(ctx context.Context, id string) (*domain.Customer, error) {
	var data domain.Customer
	if err := r.db.Model(&domain.Customer{}).
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *CustomerRepository) Update(ctx context.Context, id string, req domain.Map) (*domain.Customer, error) {
	data := &domain.Customer{}
	err := r.db.Model(&domain.Customer{}).Where("id = ?", id).Updates(req).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CustomerRepository) Delete(ctx context.Context, id string) error {
	logrus.Info("package repository Delete() Customer function called.")
	return r.db.Model(&domain.Customer{}).Where("id = ?", id).Delete(&domain.Customer{}).Error
}
