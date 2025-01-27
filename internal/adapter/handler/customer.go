package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/core/domain"
)

// Register 		godoc
// @Summary			Register a new Customer
// @Description		create a new cusomer account
// @Accept			json
// @Produce			json
// @Param			CustomerRequest	body		domain.CustomerRequest	true	"Create customer request"
// @Success			200				{object}	domain.CustomerResponse	"Customer created"
// @Router			/customers 		[post]
func (uh *Handler) CreateCustomer(ctx *gin.Context) {
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
