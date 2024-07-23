package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/adapter/config"
	"github.com/sugaml/authserver/internal/core/port"
	"github.com/sugaml/authserver/internal/core/service"
)

type Handler struct {
	config *config.HTTP
	svc    service.IService
	token  port.TokenService
	router *gin.Engine
}

func NewHandler(config *config.HTTP, svc service.IService, token port.TokenService) *Handler {
	return &Handler{
		config: config,
		svc:    svc,
		token:  token,
		router: gin.New(),
	}
}
