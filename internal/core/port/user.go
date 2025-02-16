package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

//go:generate mockgen -source=user.go -destination=mock/user.go -package=mock

// UserRepository is an interface for interacting with user-related data
type UserRepository interface {
	// Create inserts a new user into the database
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetByID selects a user by id
	GetByID(ctx context.Context, id uint64) (*domain.User, error)
	// GetByEmail selects a user by email
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	// GetByMobileNum selects a user by email
	GetByMobileNum(ctx context.Context, mobileNum string) (*domain.User, error)
	// SetPassword selects a user by email
	SetPassword(ctx context.Context, id uint64, password string) (*domain.User, error)
	// List selects a list of users with pagination
	List(ctx context.Context, skip, limit uint64) ([]*domain.User, error)
	// Update updates a user
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	// Delete deletes a user
	Delete(ctx context.Context, id uint64) error
}

// UserService is an interface for interacting with user-related business logic
type UserService interface {
	// Register registers a new user
	Register(ctx context.Context, user *domain.RegisterRequest) (*domain.UserResponse, error)

	Login(ctx context.Context, user *domain.LoginRequest) (*domain.UserResponse, error)
	// Get returns a user by id
	Get(ctx context.Context, id uint64) (*domain.UserResponse, error)
	// List returns a list of users with pagination
	List(ctx context.Context, skip, limit uint64) ([]*domain.UserResponse, error)
	// Update updates a user
	Update(ctx context.Context, user *domain.User) (*domain.UserResponse, error)
	// Delet deletes a user
	Delete(ctx context.Context, id uint64) error
}
