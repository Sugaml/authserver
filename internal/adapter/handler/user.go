package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/authserver/internal/core/domain"
)

// Register godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account with default role "cashier"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		domain.RegisterRequest	true	"Register request"
//	@Success		200				{object}	domain.UserResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/users [post]
func (uh *Handler) Register(ctx *gin.Context) {
	var req domain.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	user := domain.User{
		UserName: req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := uh.svc.User().Register(ctx, &user)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, user.NewUserResponse())
}

// listUsersRequest represents the request body for listing users
type listUsersRequest struct {
	Skip  uint64 `form:"skip" example:"0"`
	Limit uint64 `form:"limit" example:"5"`
}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	List users with pagination
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			skip	query		uint64			true	"Skip"
//	@Param			limit	query		uint64			true	"Limit"
//	@Success		200		{object}	meta			"Users displayed"
//	@Failure		400		{object}	errorResponse	"Validation error"
//	@Failure		500		{object}	errorResponse	"Internal server error"
//	@Router			/users [get]
//	@Security		BearerAuth
func (uh *Handler) ListUsers(ctx *gin.Context) {
	var req listUsersRequest
	var usersList []domain.UserResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	users, err := uh.svc.User().List(ctx, req.Skip, req.Limit)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	for _, user := range users {
		usersList = append(usersList, user.NewUserResponse())
	}

	SuccessResponse(ctx, usersList)
}

// getUserRequest represents the request body for getting a user
type getUserRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// GetUser godoc
//
//	@Summary		Get a user
//	@Description	Get a user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64			true	"User ID"
//	@Success		200	{object}	domain.UserResponse	"User displayed"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/users/{id} [get]
//	@Security		BearerAuth
func (uh *Handler) GetUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := uh.svc.User().Get(ctx, req.ID)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, user.NewUserResponse())
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update a user's name, email, password, or role by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id					path		uint64				true	"User ID"
//	@Param			updateUserRequest	body		domain.UpdateUserRequest	true	"Update user request"
//	@Success		200					{object}	domain.UserResponse		"User updated"
//	@Failure		400					{object}	errorResponse		"Validation error"
//	@Failure		401					{object}	errorResponse		"Unauthorized error"
//	@Failure		403					{object}	errorResponse		"Forbidden error"
//	@Failure		404					{object}	errorResponse		"Data not found error"
//	@Failure		500					{object}	errorResponse		"Internal server error"
//	@Router			/users/{id} [put]
//	@Security		BearerAuth
func (uh *Handler) UpdateUser(ctx *gin.Context) {
	var req domain.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	user := domain.User{
		Email:    req.Email,
		Password: req.Password,
	}
	SuccessResponse(ctx, user.NewUserResponse())
}

// deleteUserRequest represents the request body for deleting a user
type deleteUserRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64			true	"User ID"
//	@Success		200	{object}	response		"User deleted"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/users/{id} [delete]
//	@Security		BearerAuth
func (uh *Handler) DeleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err := uh.svc.User().Delete(ctx, req.ID)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, nil)
}
