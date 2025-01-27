package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/core/domain"
)

// Register godoc
//
//	@Summary		Register a new Customer
//	@Description	create a new cusomer account
//	@Accept			json
//	@Produce		json
//	@Param			CustomerRequest	body		domain.CustomerRequest	true	"Create customer request"
//	@Success		200				{object}	domain.CustomerResponse	"Customer created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/users [post]
func (uh *Handler) createCustomer(ctx *gin.Context) {
	var req domain.CustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	customer := domain.Customer{
		Username:             req.Username,
		PasswordHash:         "",
		SecurityStamp:        "",
		ConcurrencyStamp:     "",
		Email:                req.Email,
		EmailConfirmed:       false,
		PhoneNumber:          "",
		PhoneNumberConfirmed: false,
		TwoFactorEnabled:     false,
		LockoutEnabled:       false,
		AccessFailedCount:    0,
	}
	_, err := uh.svc.Customer().Create(ctx, &customer)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, customer.CustomerResponse())
}
