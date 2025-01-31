package domain

type Tenant struct {
	BaseModel
	Name               string
	Description        string
	RoleBundlesTenants []RoleBundlesTenant `gorm:"foreignkey:TenantID"`
}
