package domain

import "time"

type Tenant struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Enabled     bool   `json:"enabled"`
}

type TenantRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Enabled     bool   `json:"enabled"`
}

type TenantUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Enabled     bool   `json:"enabled"`
}

type TenantListRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Enabled     bool   `json:"enabled"`
}

func (a *Tenant) New(r *TenantRequest) *Tenant {
	a.Name = r.Name
	a.Description = r.Description
	a.Code = r.Code
	a.Enabled = r.Enabled
	return a
}

func (a *TenantUpdateRequest) NewUpdate() Map {
	return map[string]interface{}{}

}

func (a *Tenant) Validate() error {
	return nil
}

type TenantResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Code        string    `json:"code"`
	Enabled     bool      `json:"enabled"`
}
