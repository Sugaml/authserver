package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/authserver/internal/core/domain"
)

// Create a new Role
func (s *Service) CreateRole(ctx context.Context, req *domain.RoleRequest) (*domain.RoleResponse, error) {
	logrus.Info("package service Create() Role function called.")
	data := domain.Convert[domain.RoleRequest, domain.Role](req)
	data.New(req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.Role().Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("encountered %v error create Role", err)
	}
	return domain.Convert[domain.Role, domain.RoleResponse](result), nil
}

// Get returns a Role by id
func (s *Service) GetRole(ctx context.Context, id string) (*domain.RoleResponse, error) {
	logrus.Info("package service Get() Role function called.")
	result, err := s.repo.Role().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.Role, domain.RoleResponse](result), nil
}

// List returns a list of Roles with pagination
func (s *Service) ListRole(ctx context.Context, req *domain.RoleListRequest) ([]*domain.RoleResponse, int, error) {
	logrus.Info("package service List() Role function called.")
	var datas []*domain.RoleResponse
	results, count, err := s.repo.Role().List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Convert[domain.Role, domain.RoleResponse](result))
	}
	return datas, count, nil
}

// Update updates a Role
func (cs *Service) UpdateRole(ctx context.Context, id string, req *domain.RoleUpdateRequest) (*domain.RoleResponse, error) {
	_, err := cs.repo.Role().Get(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	mp := req.NewUpdate()
	result, err := cs.repo.Role().Update(ctx, id, mp)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return domain.Convert[domain.Role, domain.RoleResponse](result), nil
}

// Delet deletes a Role
func (s *Service) DeleteRole(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required Role id")
	}
	err := s.repo.Role().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
