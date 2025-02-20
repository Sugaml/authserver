package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

// type CustomerRepository interface is an interface for interacting with type Customer-related data
type CustomerRepository interface {
	Create(ctx context.Context, data *domain.Customer) (*domain.Customer, error)
	List(ctx context.Context, req *domain.ListCustomerRequest) ([]*domain.Customer, int, error)
	Get(ctx context.Context, id string) (*domain.Customer, error)
	Update(ctx context.Context, id string, req domain.Map) (*domain.Customer, error)
	Delete(ctx context.Context, id string) error
}

// type CustomerService interface is an interface for interacting with type Customer-related data
type CustomerService interface {
	CreateCustomer(ctx context.Context, data *domain.CustomerRequest) (*domain.CustomerResponse, error)
	ListCustomer(ctx context.Context, req *domain.ListCustomerRequest) ([]*domain.CustomerResponse, int, error)
	GetCustomer(ctx context.Context, id string) (*domain.CustomerResponse, error)
	UpdateCustomer(ctx context.Context, id string, req *domain.CustomerUpdateRequest) (*domain.CustomerResponse, error)
	DeleteCustomer(ctx context.Context, id string) error
}
