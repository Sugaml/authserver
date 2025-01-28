package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

// type ApplicationRepository interface is an interface for interacting with type Application-related data
type ApplicationRepository interface {
	Create(ctx context.Context, data *domain.Application) (*domain.Application, error)
	List(ctx context.Context, req *domain.ListApplicationRequest) ([]*domain.Application, int, error)
	Get(ctx context.Context, id string) (*domain.Application, error)
	Update(ctx context.Context, id string, req domain.Map) (*domain.Application, error)
	Delete(ctx context.Context, id string) error
}

// type ApplicationService interface is an interface for interacting with type Application-related data
type ApplicationService interface {
	Create(ctx context.Context, data *domain.ApplicationRequest) (*domain.ApplicationResponse, error)
	List(ctx context.Context, req *domain.ListApplicationRequest) ([]*domain.ApplicationResponse, int, error)
	Get(ctx context.Context, id string) (*domain.ApplicationResponse, error)
	Update(ctx context.Context, id string, req *domain.ApplicationUpdateRequest) (*domain.ApplicationResponse, error)
	Delete(ctx context.Context, id string) error
}
