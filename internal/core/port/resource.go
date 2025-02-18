package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

// type ResourceRepository interface is an interface for interacting with type Resource-related data
type ResourceRepository interface {
	Create(ctx context.Context, data *domain.Resource) (*domain.Resource, error)
	List(ctx context.Context, req *domain.ResourceListRequest) ([]*domain.Resource, int, error)
	Get(ctx context.Context, id string) (*domain.Resource, error)
	Update(ctx context.Context, id string, req domain.Map) (*domain.Resource, error)
	Delete(ctx context.Context, id string) error
}

// type ResourceService interface is an interface for interacting with type Resource-related data
type ResourceService interface {
	CreateResource(ctx context.Context, data *domain.ResourceRequest) (*domain.ResourceResponse, error)
	ListResource(ctx context.Context, req *domain.ResourceListRequest) ([]*domain.ResourceResponse, int, error)
	GetResource(ctx context.Context, id string) (*domain.ResourceResponse, error)
	UpdateResource(ctx context.Context, id string, req *domain.ResourceUpdateRequest) (*domain.ResourceResponse, error)
	DeleteResource(ctx context.Context, id string) error
}
