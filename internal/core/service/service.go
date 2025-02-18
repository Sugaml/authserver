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

func NewService(repo repository.IRepository) port.IService {
	return &Service{
		repo: repo,
	}
}

func GetOauthServer(repo repository.IRepository) *server.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// Token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := newClientStotreService(repo)
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
