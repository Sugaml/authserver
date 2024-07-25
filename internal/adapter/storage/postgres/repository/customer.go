package repository

import (
	"context"

	"github.com/jinzhu/gorm"
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

// Create creates a new Customer in the database
func (r *CustomerRepository) Create(ctx context.Context, data *domain.Customer) (*domain.Customer, error) {
	err := r.db.Model(&domain.Customer{}).Create(data).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

// GetByID gets a customer by ID from the database
func (r *CustomerRepository) GetByID(ctx context.Context, id uint64) (*domain.Customer, error) {
	data := &domain.Customer{}
	err := r.db.Model(domain.Customer{}).Where("id = ? and is_active = true", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

// List lists all customers from the database
func (r *CustomerRepository) List(ctx context.Context, skip, limit uint64) ([]domain.Customer, error) {
	customers := []domain.Customer{}
	err := r.db.Model(&domain.Customer{}).Order("id desc").Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, err
}

// Update updates a customer by ID in the database
func (r *CustomerRepository) Update(ctx context.Context, data *domain.Customer) (*domain.Customer, error) {
	customer := &domain.Customer{}
	if data.Email != "" {
		customer.Email = data.Email
	}
	err := r.db.Model(&domain.Customer{}).Where("id = ?", data.ID).Updates(customer).Error
	if err != nil {
		return &domain.Customer{}, err
	}
	return customer, nil
}

// Delete deletes a customer by ID from the database
func (r *CustomerRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.Model(&domain.Customer{}).Where("id = ?", id).Delete(&domain.Customer{}).Error
}
