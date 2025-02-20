package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/core/domain"
)

// AddClientSecret	godoc
// @Summary			Add a new ClientSecret
// @Description		Add a new ClientSecret
// @Tags			ClientSecret
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			ClientSecretRequest			body		domain.ClientSecretRequest		true		"Add ClientSecret Request"
// @Success			200							{array}		domain.ClientSecretResponse					"Business  ClientSecret created"
// @Router			/client-secret 		[post]
func (ph *Handler) CreateClientSecret(ctx *gin.Context) {
	var req *domain.ClientSecretRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := ph.svc.CreateSecret(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// GetClientSecret 	godoc
// @Summary 		Get ClientSecret
// @Description 	Get ClientSecret from Id
// @Tags 			ClientSecret
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "ClientSecret id"
// @Success 		200 {object} domain.ClientSecretResponse
// @Router 			/client-secret/{id} [get]
func (ch *Handler) GetClientSecret(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := ch.svc.GetSecret(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateClientSecret	godoc
// @Summary 			Update ClientSecret
// @Description 		Update ClientSecret from Id
// @Tags 				ClientSecret
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 								path 		string 								true 	"ClientSecret id"
// @Param 				UpdateClientSecretRequest	 	body 		domain.ClientSecretUpdateRequest 	true 	"Update ClientSecret Response request"
// @Success 			200 							{object} 	domain.ClientSecretResponse
// @Router 				/client-secret/{id} 				[put]
func (h *Handler) UpdateClientSecret(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.ClientSecretUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	_, err := h.svc.UpdateSecret(ctx, id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	result, err := h.svc.GetSecret(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// DeleteClientSecret 	godoc
// @Summary 			Delete ClientSecret
// @Description 		Delete ClientSecret from Id
// @Tags 				ClientSecret
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 						path 		string 						true 	"ClientSecret id"
// @Success 			200 					{object} 	domain.ClientSecretResponse
// @Router 				/client-secret/{id} 	[delete]
func (ch *Handler) DeleteClientSecret(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required ClientSecret id"))
		return
	}
	result, err := ch.svc.GetSecret(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = ch.svc.DeleteSecret(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}
