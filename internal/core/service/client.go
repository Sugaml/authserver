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

type ClientServiceGetter interface {
	Client() port.ClientService
}

type ClientService struct {
	repo repository.IRepository
}

func newClientService(repo repository.IRepository) *ClientService {
	return &ClientService{
		repo: repo,
	}
}

// Create a new client
func (s *ClientService) Create(ctx context.Context, req *domain.ClientRequest) (*domain.ClientResponse, error) {
	logrus.Info("package service Create() client function called.")
	data := &domain.Client{}
	data.New(req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.Client().Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("encountered %v error create client", err)
	}
	return domain.Response[domain.Client, domain.ClientResponse](result), nil
}

// Get returns a client by id
func (s *ClientService) Get(ctx context.Context, id string) (*domain.ClientResponse, error) {
	logrus.Info("package service Get() client function called.")
	result, err := s.repo.Client().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Response[domain.Client, domain.ClientResponse](result), nil
}

func (s *ClientService) ListByApplicationID(ctx context.Context, id string, req *domain.ClientListRequest) ([]*domain.ClientResponse, int, error) {
	logrus.Info("package service ListByApplicationID() client function called.")
	var datas []*domain.ClientResponse
	results, count, err := s.repo.Client().ListByApplicationID(ctx, id, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Response[domain.Client, domain.ClientResponse](result))
	}
	return datas, count, nil
}

// List returns a list of clients with pagination
func (s *ClientService) List(ctx context.Context, req *domain.ClientListRequest) ([]*domain.ClientResponse, int, error) {
	logrus.Info("package service List() client function called.")
	var datas []*domain.ClientResponse
	results, count, err := s.repo.Client().List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Response[domain.Client, domain.ClientResponse](result))
	}
	return datas, count, nil
}

// Update updates a client
func (cs *ClientService) Update(ctx context.Context, id string, req *domain.ClientUpdateRequest) (*domain.ClientResponse, error) {
	_, err := cs.repo.Client().Get(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	mp := req.NewUpdate()
	result, err := cs.repo.Client().Update(ctx, id, mp)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return domain.Response[domain.Client, domain.ClientResponse](result), nil
}

// Delet deletes a client
func (s *ClientService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required client id")
	}
	err := s.repo.Client().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
