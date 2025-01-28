package service

import (
	"log"

	"github.com/sugaml/authserver/internal/adapter/storage/postgres/repository"
	"github.com/sugaml/authserver/internal/core/port"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type Service struct {
	repo repository.IRepository
}

type IService interface {
	GetOauthServer() *server.Server
	UserServiceGetter
	CustomerServiceGetter
	ClientServiceGetter
	ClientSecretServiceGetter
}

func NewService(repo repository.IRepository) IService {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetOauthServer() *server.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// Token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := newClientStotreService(s.repo)
	// Client database store
	manager.MapClientStorage(clientStore)

	// Set custom JWT generator
	manager.MapAccessGenerate(&JWTAccessGenerate{})

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
	return srv
}

func (s *Service) User() port.UserService {
	return newUserService(s.repo)
}

func (s *Service) Customer() port.CustomerService {
	return newCustomerService(s.repo)
}

func (s *Service) Client() port.ClientService {
	return newClientService(s.repo)
}

func (s *Service) ClientSecret() port.ClientSecretService {
	return newClientSecretService(s.repo)
}
