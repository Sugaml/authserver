package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/authserver/internal/core/domain"
)

// Create a new Tenant
func (s *Service) CreateTenant(ctx context.Context, req *domain.TenantRequest) (*domain.TenantResponse, error) {
	logrus.Info("package service Create() Tenant function called.")
	data := domain.Convert[domain.TenantRequest, domain.Tenant](req)
	data.New(req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.Tenant().Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("encountered %v error create Tenant", err)
	}
	return domain.Convert[domain.Tenant, domain.TenantResponse](result), nil
}

// Get returns a Tenant by id
func (s *Service) GetTenant(ctx context.Context, id string) (*domain.TenantResponse, error) {
	logrus.Info("package service Get() Tenant function called.")
	result, err := s.repo.Tenant().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.Tenant, domain.TenantResponse](result), nil
}

// List returns a list of Tenants with pagination
func (s *Service) ListTenant(ctx context.Context, req *domain.TenantListRequest) ([]*domain.TenantResponse, int, error) {
	logrus.Info("package service List() Tenant function called.")
	var datas []*domain.TenantResponse
	results, count, err := s.repo.Tenant().List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Convert[domain.Tenant, domain.TenantResponse](result))
	}
	return datas, count, nil
}

// Update updates a Tenant
func (cs *Service) UpdateTenant(ctx context.Context, id string, req *domain.TenantUpdateRequest) (*domain.TenantResponse, error) {
	_, err := cs.repo.Tenant().Get(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	mp := req.NewUpdate()
	result, err := cs.repo.Tenant().Update(ctx, id, mp)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return domain.Convert[domain.Tenant, domain.TenantResponse](result), nil
}

// Delet deletes a Tenant
func (s *Service) DeleteTenant(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required tenant id")
	}
	err := s.repo.Tenant().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
