package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/core/domain"
)

// AddCustomer		godoc
// @Summary			Add a new Customer
// @Description		Add a new Customer
// @Tags			Customer
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			CustomerRequest			body		domain.CustomerRequest		true		"Add Customer Request"
// @Success			200						{array}		domain.CustomerResponse					"Customer created"
// @Router			/customer 				[post]
func (p *Handler) CreateCustomer(ctx *gin.Context) {
	var req *domain.CustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := p.svc.Customer().Create(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListCustomer 			godoc
// @Summary 					List customer
// @Description 				List customer from provider ID
// @Tags 						Customer
// @Accept  					json
// @Produce  					json
// @Security 					ApiKeyAuth
// @Param 						id 				path 		string 		true 		"Provider id"
// @Success 					200 			{array} 	domain.CustomerResponse
// @Router 						/customer	 	[get]
func (h *Handler) ListProviderCustomer(ctx *gin.Context) {
	var req domain.ListCustomerRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, 400, err)
		return
	}
	result, count, err := h.svc.Customer().List(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, int(req.Page), int(req.Size)))
}

// GetCustomer 	godoc
// @Summary 		Get Customer
// @Description 	Get Customer from Id
// @Tags 			Customer
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "Customer id"
// @Success 		200 {object} domain.CustomerResponse
// @Router 			/customer/{id} [get]
func (h *Handler) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.Customer().Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateCustomer		godoc
// @Summary 			Update Customer
// @Description 		Update Customer from Id
// @Tags 				Customer
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 								path 			string 								true 	"Customer id"
// @Param 				CustomerUpdateRequest	 		body 			domain.CustomerUpdateRequest 	true 	"Update Customer Response request"
// @Success 			200 							{object} 		domain.CustomerResponse
// @Router 				/customer/{id} 					[put]
func (h *Handler) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.CustomerUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	_, err := h.svc.Customer().Update(ctx, id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	result, err := h.svc.Customer().Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}

// DeleteCustomer 		godoc
// @Summary 			Delete Customer
// @Description 		Delete Customer from Id
// @Tags 				Customer
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 						path 		string 						true 	"Customer id"
// @Success 			200 					{object} 	domain.CustomerResponse
// @Router 				/customer/{id} 			[delete]
func (ch *Handler) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required customer id"))
		return
	}
	result, err := ch.svc.Customer().Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = ch.svc.Customer().Delete(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result)
}
