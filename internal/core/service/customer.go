package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/authserver/internal/core/domain"
)

// Create a new Customer
func (s *Service) CreateCustomer(ctx context.Context, req *domain.CustomerRequest) (*domain.CustomerResponse, error) {
	logrus.Info("package service Create() Customer function called.")
	data := domain.Convert[domain.CustomerRequest, domain.Customer](req)
	data.New(req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.Customer().Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("encountered %v error create Customer", err)
	}
	return domain.Convert[domain.Customer, domain.CustomerResponse](result), nil
}

// Get returns a Customer by id
func (s *Service) GetCustomer(ctx context.Context, id string) (*domain.CustomerResponse, error) {
	logrus.Info("package service Get() Customer function called.")
	result, err := s.repo.Customer().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.Customer, domain.CustomerResponse](result), nil
}

// List returns a list of Customers with pagination
func (s *Service) ListCustomer(ctx context.Context, req *domain.ListCustomerRequest) ([]*domain.CustomerResponse, int, error) {
	logrus.Info("package service List() Customer function called.")
	var datas []*domain.CustomerResponse
	results, count, err := s.repo.Customer().List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Convert[domain.Customer, domain.CustomerResponse](result))
	}
	return datas, count, nil
}

// Update updates a Customer
func (cs *Service) UpdateCustomer(ctx context.Context, id string, req *domain.CustomerUpdateRequest) (*domain.CustomerResponse, error) {
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

	return domain.Convert[domain.Customer, domain.CustomerResponse](result), nil
}

// Delet deletes a Customer
func (s *Service) DeleteCustomer(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required customer id")
	}
	err := s.repo.Customer().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
