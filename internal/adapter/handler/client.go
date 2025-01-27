package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/core/domain"
)

// AddClient	godoc
// @Summary			Add a new Client
// @Description		Add a new Client
// @Tags			Client
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			ClientRequest			body		domain.ClientRequest		true		"Add Client Request"
// @Success			200							{array}		domain.ClientResponse					"Business  Client created"
// @Router			/client 		[post]
func (ph *Handler) CreateClient(ctx *gin.Context) {
	var req *domain.ClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := ph.svc.Client().Create(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// GetClient 		godoc
// @Summary 		Get Client
// @Description 	Get Client from Id
// @Tags 			Client
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "Client id"
// @Success 		200 {object} domain.ClientResponse
// @Router 			/client/{id} [get]
func (ch *Handler) GetClient(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := ch.svc.Client().Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateClient	godoc
// @Summary 			Update Client
// @Description 		Update Client from Id
// @Tags 				Client
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 								path 		string 								true 	"Client id"
// @Param 				UpdateClientRequest	 			body 		domain.ClientUpdateRequest 	true 	"Update Client Response request"
// @Success 			200 							{object} 	domain.ClientResponse
// @Router 				/client/{id} 				[put]
func (h *Handler) UpdateClient(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.ClientUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := h.svc.Client().Update(ctx, id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	result, err := h.svc.Client().Get(ctx, data.ID)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// DeleteClient 		godoc
// @Summary 			Delete Client
// @Description 		Delete Client from Id
// @Tags 				Client
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 						path 		string 						true 	"Client id"
// @Success 			200 					{object} 	domain.ClientResponse
// @Router 				/client/{id} 	[delete]
func (ch *Handler) DeleteClient(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required client id"))
		return
	}
	result, err := ch.svc.Client().Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = ch.svc.Client().Delete(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}
