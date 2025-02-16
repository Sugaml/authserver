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

type ClientSecretServiceGetter interface {
	ClientSecret() port.ClientSecretService
}

type SecretService struct {
	repo repository.IRepository
}

func newClientSecretService(repo repository.IRepository) *SecretService {
	return &SecretService{
		repo: repo,
	}
}

// Create a new Secret
func (s *SecretService) Create(ctx context.Context, req *domain.ClientSecretRequest) (*domain.ClientSecretResponse, error) {
	logrus.Info("package service Create() Secret function called.")
	data := domain.Convert[domain.ClientSecretRequest, domain.ClientSecret](req)
	data.New(req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.ClientSecret().Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("encountered %v error create Secret", err)
	}
	return domain.Convert[domain.ClientSecret, domain.ClientSecretResponse](result), nil
}

// Get returns a Secret by id
func (s *SecretService) Get(ctx context.Context, id string) (*domain.ClientSecretResponse, error) {
	logrus.Info("package service Get() Secret function called.")
	result, err := s.repo.ClientSecret().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.ClientSecret, domain.ClientSecretResponse](result), nil
}

func (s *SecretService) ListByApplicationID(ctx context.Context, id string, req *domain.ClientSecretListRequest) ([]*domain.ClientSecretResponse, int, error) {
	logrus.Info("package service ListByApplicationID() Secret function called.")
	var datas []*domain.ClientSecretResponse
	results, count, err := s.repo.ClientSecret().ListByApplicationID(ctx, id, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Convert[domain.ClientSecret, domain.ClientSecretResponse](result))
	}
	return datas, count, nil
}

// List returns a list of Secrets with pagination
func (s *SecretService) List(ctx context.Context, req *domain.ClientSecretListRequest) ([]*domain.ClientSecretResponse, int, error) {
	logrus.Info("package service List() Secret function called.")
	var datas []*domain.ClientSecretResponse
	results, count, err := s.repo.ClientSecret().List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Convert[domain.ClientSecret, domain.ClientSecretResponse](result), nil)
	}
	return datas, count, nil
}

// Update updates a Secret
func (cs *SecretService) Update(ctx context.Context, id string, req *domain.ClientSecretUpdateRequest) (*domain.ClientSecretResponse, error) {
	_, err := cs.repo.ClientSecret().Get(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	mp := req.NewUpdate()
	result, err := cs.repo.ClientSecret().Update(ctx, id, mp)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return domain.Convert[domain.ClientSecret, domain.ClientSecretResponse](result), nil
}

// Delet deletes a Secret
func (s *SecretService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("required client secret id")
	}
	err := s.repo.ClientSecret().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
