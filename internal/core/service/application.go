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

type ApplicationServiceGetter interface {
	Application() port.ApplicationService
}

type ApplicationService struct {
	repo repository.IRepository
}

func newApplicationService(repo repository.IRepository) *ApplicationService {
	return &ApplicationService{
		repo: repo,
	}
}

// Create a new Application
func (s *ApplicationService) Create(ctx context.Context, req *domain.ApplicationRequest) (*domain.ApplicationResponse, error) {
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
func (s *ApplicationService) Get(ctx context.Context, id string) (*domain.ApplicationResponse, error) {
	logrus.Info("package service Get() Application function called.")
	result, err := s.repo.Application().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.Application, domain.ApplicationResponse](result), nil
}

// List returns a list of Applications with pagination
func (s *ApplicationService) List(ctx context.Context, req *domain.ListApplicationRequest) ([]*domain.ApplicationResponse, int, error) {
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
func (cs *ApplicationService) Update(ctx context.Context, id string, req *domain.ApplicationUpdateRequest) (*domain.ApplicationResponse, error) {
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
func (s *ApplicationService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required Application id")
	}
	err := s.repo.Application().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
