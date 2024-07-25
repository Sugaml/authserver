package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

//go:generate mockgen -source=customer.go -destination=mock/customer.go -package=mock

// CustomerRepository is an interface for interacting with ustomer-related data
type CustomerRepository interface {
	// Create inserts a new customer into the database
	Create(ctx context.Context, data *domain.Customer) (*domain.Customer, error)
	// GetByID selects a customer by id
	GetByID(ctx context.Context, id uint64) (*domain.Customer, error)

	// List selects a list of customers with pagination
	List(ctx context.Context, skip, limit uint64) ([]domain.Customer, error)
	// Update updates a customer
	Update(ctx context.Context, user *domain.Customer) (*domain.Customer, error)
	// Delete deletes a customer
	Delete(ctx context.Context, id uint64) error
}

// CustomerService is an interface for interacting with customer-related business logic
type CustomerService interface {
	// Create a new customer
	Create(ctx context.Context, data *domain.Customer) (*domain.Customer, error)
	// Get returns a customer by id
	Get(ctx context.Context, id uint64) (*domain.Customer, error)
	// List returns a list of customers with pagination
	List(ctx context.Context, skip, limit uint64) ([]domain.Customer, error)
	// Update updates a customer
	Update(ctx context.Context, customer *domain.Customer) (*domain.Customer, error)
	// Delet deletes a customer
	Delete(ctx context.Context, id uint64) error
}
