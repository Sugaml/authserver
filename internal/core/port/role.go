package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

// type RoleRepository interface is an interface for interacting with type Role-related data
type RoleRepository interface {
	Create(ctx context.Context, data *domain.Role) (*domain.Role, error)
	List(ctx context.Context, req *domain.RoleListRequest) ([]*domain.Role, int, error)
	Get(ctx context.Context, id string) (*domain.Role, error)
	Update(ctx context.Context, id string, req domain.Map) (*domain.Role, error)
	Delete(ctx context.Context, id string) error
}

// type RoleService interface is an interface for interacting with type Role-related data
type RoleService interface {
	Create(ctx context.Context, data *domain.RoleRequest) (*domain.RoleResponse, error)
	List(ctx context.Context, req *domain.RoleListRequest) ([]*domain.RoleResponse, int, error)
	Get(ctx context.Context, id string) (*domain.RoleResponse, error)
	Update(ctx context.Context, id string, req *domain.RoleUpdateRequest) (*domain.RoleResponse, error)
	Delete(ctx context.Context, id string) error
}
