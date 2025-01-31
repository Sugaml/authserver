package domain

import "time"

type RoleClaim struct {
	BaseModel
	RoleID     string
	ClaimType  string
	ClaimValue string
}

type Role struct {
	BaseModel
	Name             string
	NormalizedName   string
	RoleClaims       []RoleClaim       `gorm:"foreignkey:RoleID"`
	RoleBundlesRoles []RoleBundlesRole `gorm:"foreignkey:RoleID"`
}

type RoleRequest struct {
	Name           string `json:"name"`
	NormalizedName string `json:"normalized_name"`
}

type RoleUpdateRequest struct {
	Name           string `json:"name"`
	NormalizedName string `json:"normalized_name"`
}

type RoleResponse struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
	NormalizedName string    `json:"normalized_name"`
}

type RoleBundle struct {
	BaseModel
	Name             string
	Description      string
	RoleBundlesRoles []RoleBundlesRole `gorm:"foreignkey:RoleBundleID"`
}

type RoleBundlesRole struct {
	BaseModel
	RoleBundleID string
	RoleID       string
}

func (a *Role) New(r *RoleRequest) {
	a.Name = r.Name
	a.NormalizedName = r.NormalizedName
}

func (a *Role) Validate() error {
	return nil
}

func (r *RoleUpdateRequest) NewUpdate() Map {
	return map[string]interface{}{}
}
