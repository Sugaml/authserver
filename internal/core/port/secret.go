package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

// type ClientSecretRepository interface is an interface for interacting with type Announcement-related data
type ClientSecretRepository interface {
	Create(ctx context.Context, data *domain.ClientSecret) (*domain.ClientSecret, error)
	List(ctx context.Context, req *domain.ClientSecretListRequest) ([]*domain.ClientSecret, int, error)
	ListByApplicationID(ctx context.Context, id string, req *domain.ClientSecretListRequest) ([]*domain.ClientSecret, int, error)
	Get(ctx context.Context, id string) (*domain.ClientSecret, error)
	Update(ctx context.Context, id string, req domain.Map) (*domain.ClientSecret, error)
	UpdateIsActive(ctx context.Context, id string, isActive bool) (*domain.ClientSecret, error)
	Delete(ctx context.Context, id string) error
}

// type ClientSecretService interface is an interface for interacting with type Announcement-related data
type ClientSecretService interface {
	Create(ctx context.Context, data *domain.ClientSecretRequest) (*domain.ClientSecretResponse, error)
	List(ctx context.Context, req *domain.ClientSecretListRequest) ([]*domain.ClientSecretResponse, int, error)
	ListByApplicationID(ctx context.Context, id string, req *domain.ClientSecretListRequest) ([]*domain.ClientSecretResponse, int, error)
	Get(ctx context.Context, id string) (*domain.ClientSecretResponse, error)
	Update(ctx context.Context, id string, req *domain.ClientSecretUpdateRequest) (*domain.ClientSecretResponse, error)
	Delete(ctx context.Context, id string) error
}
