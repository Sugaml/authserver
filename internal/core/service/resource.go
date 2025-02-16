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

type ResourceServiceGetter interface {
	Resource() port.ResourceService
}

type ResourceService struct {
	repo repository.IRepository
}

func newResourceService(repo repository.IRepository) *ResourceService {
	return &ResourceService{
		repo: repo,
	}
}

// Create a new Resource
func (s *ResourceService) Create(ctx context.Context, req *domain.ResourceRequest) (*domain.ResourceResponse, error) {
	logrus.Info("package service Create() Resource function called.")
	data := domain.Convert[domain.ResourceRequest, domain.Resource](req)
	data.New(req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.Resource().Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("encountered %v error create Resource", err)
	}
	return domain.Convert[domain.Resource, domain.ResourceResponse](result), nil
}

// Get returns a Resource by id
func (s *ResourceService) Get(ctx context.Context, id string) (*domain.ResourceResponse, error) {
	logrus.Info("package service Get() Resource function called.")
	result, err := s.repo.Resource().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.Resource, domain.ResourceResponse](result), nil
}

// List returns a list of Resources with pagination
func (s *ResourceService) List(ctx context.Context, req *domain.ResourceListRequest) ([]*domain.ResourceResponse, int, error) {
	logrus.Info("package service List() Resource function called.")
	var datas []*domain.ResourceResponse
	results, count, err := s.repo.Resource().List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Convert[domain.Resource, domain.ResourceResponse](result))
	}
	return datas, count, nil
}

// Update updates a Resource
func (cs *ResourceService) Update(ctx context.Context, id string, req *domain.ResourceUpdateRequest) (*domain.ResourceResponse, error) {
	_, err := cs.repo.Resource().Get(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	mp := req.NewUpdate()
	result, err := cs.repo.Resource().Update(ctx, id, mp)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return domain.Convert[domain.Resource, domain.ResourceResponse](result), nil
}

// Delet deletes a Resource
func (s *ResourceService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required resource id")
	}
	err := s.repo.Resource().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
