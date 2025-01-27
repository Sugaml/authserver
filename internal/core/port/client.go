package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

// type ClientRepository interface is an interface for interacting with type Announcement-related data
type ClientRepository interface {
	Create(ctx context.Context, data *domain.Client) (*domain.Client, error)
	List(ctx context.Context, req *domain.ClientListRequest) ([]*domain.Client, int, error)
	ListByApplicationID(ctx context.Context, id string, req *domain.ClientListRequest) ([]*domain.Client, int, error)
	Get(ctx context.Context, id string) (*domain.Client, error)
	Update(ctx context.Context, id string, req domain.Map) (*domain.Client, error)
	UpdateIsActive(ctx context.Context, id string, isActive bool) (*domain.Client, error)
	Delete(ctx context.Context, id string) error
}

// type ClientService interface is an interface for interacting with type Announcement-related data
type ClientService interface {
	Create(ctx context.Context, data *domain.ClientRequest) (*domain.ClientResponse, error)
	List(ctx context.Context, req *domain.ClientListRequest) ([]*domain.ClientResponse, int, error)
	ListByApplicationID(ctx context.Context, id string, req *domain.ClientListRequest) ([]*domain.ClientResponse, int, error)
	Get(ctx context.Context, id string) (*domain.ClientResponse, error)
	Update(ctx context.Context, id string, req *domain.ClientUpdateRequest) (*domain.ClientResponse, error)
	Delete(ctx context.Context, id string) error
}
