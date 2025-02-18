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
	CreateRole(ctx context.Context, data *domain.RoleRequest) (*domain.RoleResponse, error)
	ListRole(ctx context.Context, req *domain.RoleListRequest) ([]*domain.RoleResponse, int, error)
	GetRole(ctx context.Context, id string) (*domain.RoleResponse, error)
	UpdateRole(ctx context.Context, id string, req *domain.RoleUpdateRequest) (*domain.RoleResponse, error)
	DeleteRole(ctx context.Context, id string) error
}
