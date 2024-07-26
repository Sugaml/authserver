package service

import (
	"context"

	"github.com/sugaml/authserver/internal/adapter/storage/postgres/repository"
	"github.com/sugaml/authserver/internal/core/domain"
	"github.com/sugaml/authserver/internal/core/port"
)

type CustomerServiceGetter interface {
	Customer() port.CustomerService
}

type CustomerService struct {
	repo repository.IRepository
}

func newCustomerService(repo repository.IRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

// Create a new customer
func (cs *CustomerService) Create(ctx context.Context, data *domain.Customer) (*domain.Customer, error) {
	data, err := cs.repo.Customer().Create(ctx, data)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return data, nil
}

// Get returns a customer by id
func (cs *CustomerService) Get(ctx context.Context, id uint64) (*domain.Customer, error) {
	data, err := cs.repo.Customer().GetByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return data, nil
}

// List returns a list of customers with pagination
func (cs *CustomerService) List(ctx context.Context, skip, limit uint64) ([]domain.Customer, error) {
	var datas []domain.Customer

	datas, err := cs.repo.Customer().List(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return datas, nil
}

// Update updates a customer
func (cs *CustomerService) Update(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	_, err := cs.repo.Customer().GetByID(ctx, uint64(customer.ID))
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	_, err = cs.repo.Customer().Update(ctx, customer)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return customer, nil
}

// Delet deletes a customer
func (cs *CustomerService) Delete(ctx context.Context, id uint64) error {
	_, err := cs.repo.Customer().GetByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return domain.ErrInternal
	}

	return cs.repo.Customer().Delete(ctx, id)
}
