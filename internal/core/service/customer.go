package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
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

// Create a new Customer
func (s *CustomerService) Create(ctx context.Context, req *domain.CustomerRequest) (*domain.CustomerResponse, error) {
	logrus.Info("package service Create() Customer function called.")
	data := &domain.Customer{}
	data.New(req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.Customer().Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("encountered %v error create Customer", err)
	}
	return domain.Response[domain.Customer, domain.CustomerResponse](result), nil
}

// Get returns a Customer by id
func (s *CustomerService) Get(ctx context.Context, id string) (*domain.CustomerResponse, error) {
	logrus.Info("package service Get() Customer function called.")
	result, err := s.repo.Customer().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Response[domain.Customer, domain.CustomerResponse](result), nil
}

// List returns a list of Customers with pagination
func (s *CustomerService) List(ctx context.Context, req *domain.ListCustomerRequest) ([]*domain.CustomerResponse, int, error) {
	logrus.Info("package service List() Customer function called.")
	var datas []*domain.CustomerResponse
	results, count, err := s.repo.Customer().List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Response[domain.Customer, domain.CustomerResponse](result))
	}
	return datas, count, nil
}

// Update updates a Customer
func (cs *CustomerService) Update(ctx context.Context, id string, req *domain.CustomerUpdateRequest) (*domain.CustomerResponse, error) {
	_, err := cs.repo.Customer().Get(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	mp := req.NewUpdate()
	result, err := cs.repo.Customer().Update(ctx, id, mp)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return domain.Response[domain.Customer, domain.CustomerResponse](result), nil
}

// Delet deletes a Customer
func (s *CustomerService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required customer id")
	}
	err := s.repo.Customer().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
