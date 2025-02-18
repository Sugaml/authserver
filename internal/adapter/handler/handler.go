package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/adapter/config"
	"github.com/sugaml/authserver/internal/core/port"
	"gopkg.in/oauth2.v3/server"
)

type Handler struct {
	config *config.HTTP
	svc    port.IService
	token  port.TokenService
	srv    *server.Server
	router *gin.Engine
}

func NewHandler(config *config.HTTP, svc port.IService, token port.TokenService, srv *server.Server) *Handler {
	return &Handler{
		config: config,
		svc:    svc,
		token:  token,
		srv:    srv,
		router: gin.New(),
	}
}
