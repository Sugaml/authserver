package service

import (
	"github.com/sugaml/authserver/internal/adapter/storage/postgres/repository"
	"github.com/sugaml/authserver/internal/core/port"
)

type Service struct {
	repo repository.IRepository
}

type IService interface {
	UserServiceGetter
}

func NewService(repo repository.IRepository) IService {
	return &Service{
		repo: repo,
	}
}

func (s *Service) User() port.UserService {
	return newUserService(s.repo)
}
