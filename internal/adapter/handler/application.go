package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/core/domain"
)

// AddApplication	godoc
// @Summary			Add a new Application
// @Description		Add a new Application
// @Tags			Application
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			ApplicationRequest			body		domain.ApplicationRequest		true		"Add Application Request"
// @Success			200						{array}		domain.ApplicationResponse					"Application created"
// @Router			/application 				[post]
func (p *Handler) CreateApplication(ctx *gin.Context) {
	var req *domain.ApplicationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := p.svc.Application().Create(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListApplication 				godoc
// @Summary 					List Application
// @Description 				List Application from provider ID
// @Tags 						Application
// @Accept  					json
// @Produce  					json
// @Security 					ApiKeyAuth
// @Param 						id 				path 		string 		true 		"Provider id"
// @Success 					200 			{array} 	domain.ApplicationResponse
// @Router 						/application	 	[get]
func (h *Handler) ListProviderApplication(ctx *gin.Context) {
	var req domain.ListApplicationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, 400, err)
		return
	}
	result, count, err := h.svc.Application().List(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, int(req.Page), int(req.Size)))
}

// GetApplication 	godoc
// @Summary 		Get Application
// @Description 	Get Application from Id
// @Tags 			Application
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "Application id"
// @Success 		200 {object} domain.ApplicationResponse
// @Router 			/application/{id} [get]
func (h *Handler) GetApplication(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.Application().Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateApplication		godoc
// @Summary 			Update Application
// @Description 		Update Application from Id
// @Tags 				Application
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 								path 			string 								true 	"Application id"
// @Param 				ApplicationUpdateRequest	 		body 		domain.ApplicationUpdateRequest 	true 	"Update Application Response request"
// @Success 			200 							{object} 		domain.ApplicationResponse
// @Router 				/application/{id} 					[put]
func (h *Handler) UpdateApplication(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.ApplicationUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	_, err := h.svc.Application().Update(ctx, id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	result, err := h.svc.Application().Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// DeleteApplication 	godoc
// @Summary 			Delete Application
// @Description 		Delete Application from Id
// @Tags 				Application
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 						path 		string 						true 	"Application id"
// @Success 			200 					{object} 	domain.ApplicationResponse
// @Router 				/application/{id} 			[delete]
func (ch *Handler) DeleteApplication(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required Application id"))
		return
	}
	result, err := ch.svc.Application().Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = ch.svc.Application().Delete(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}
