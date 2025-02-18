package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/authserver/internal/core/domain"
)

// Create a new Application
func (s *Service) CreateApplication(ctx context.Context, req *domain.ApplicationRequest) (*domain.ApplicationResponse, error) {
	logrus.Info("package service Create() Application function called.")
	data := domain.Convert[domain.ApplicationRequest, domain.Application](req)
	data.New(req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.Application().Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("encountered %v error create Application", err)
	}
	return domain.Convert[domain.Application, domain.ApplicationResponse](result), nil
}

// Get returns a Application by id
func (s *Service) GetApplication(ctx context.Context, id string) (*domain.ApplicationResponse, error) {
	logrus.Info("package service Get() Application function called.")
	result, err := s.repo.Application().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.Application, domain.ApplicationResponse](result), nil
}

// List returns a list of Applications with pagination
func (s *Service) ListApplication(ctx context.Context, req *domain.ListApplicationRequest) ([]*domain.ApplicationResponse, int, error) {
	logrus.Info("package service List() Application function called.")
	var datas []*domain.ApplicationResponse
	results, count, err := s.repo.Application().List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Convert[domain.Application, domain.ApplicationResponse](result))
	}
	return datas, count, nil
}

// Update updates a Application
func (cs *Service) UpdateApplication(ctx context.Context, id string, req *domain.ApplicationUpdateRequest) (*domain.ApplicationResponse, error) {
	_, err := cs.repo.Application().Get(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	mp := req.NewUpdate()
	result, err := cs.repo.Application().Update(ctx, id, mp)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return domain.Convert[domain.Application, domain.ApplicationResponse](result), nil
}

// Delet deletes a Application
func (s *Service) DeleteApplication(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required application id")
	}
	err := s.repo.Application().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
