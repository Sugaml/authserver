package port

import (
	"context"

	"github.com/sugaml/authserver/internal/core/domain"
)

// type TenantRepository interface is an interface for interacting with type Tenant-related data
type TenantRepository interface {
	Create(ctx context.Context, data *domain.Tenant) (*domain.Tenant, error)
	List(ctx context.Context, req *domain.TenantListRequest) ([]*domain.Tenant, int, error)
	Get(ctx context.Context, id string) (*domain.Tenant, error)
	Update(ctx context.Context, id string, req domain.Map) (*domain.Tenant, error)
	Delete(ctx context.Context, id string) error
}

// type TenantService interface is an interface for interacting with type Tenant-related data
type TenantService interface {
	CreateTenant(ctx context.Context, data *domain.TenantRequest) (*domain.TenantResponse, error)
	ListTenant(ctx context.Context, req *domain.TenantListRequest) ([]*domain.TenantResponse, int, error)
	GetTenant(ctx context.Context, id string) (*domain.TenantResponse, error)
	UpdateTenant(ctx context.Context, id string, req *domain.TenantUpdateRequest) (*domain.TenantResponse, error)
	DeleteTenant(ctx context.Context, id string) error
}
