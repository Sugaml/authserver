package domain

import "time"

type Application struct {
	BaseModel
	Name     string `json:"Name"`
	Logo     string `json:"logo"`
	Owner    string `json:"owner"`
	IsActive bool   `json:"is_active"`
}

type ApplicationRequest struct {
	Name     string `json:"Name"`
	Logo     string `json:"logo"`
	Owner    string `json:"owner"`
	IsActive bool   `json:"is_active"`
}

type ApplicationUpdateRequest struct {
	Name     string `json:"Name"`
	Logo     string `json:"logo"`
	Owner    string `json:"owner"`
	IsActive *bool  `json:"is_active"`
}

type ListApplicationRequest struct {
	ListRequest
	Name     string `json:"Name"`
	Logo     string `json:"logo"`
	Owner    string `json:"owner"`
	IsActive bool   `json:"is_active"`
}

type ApplicationResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"Name"`
	Logo      string    `json:"logo"`
	Owner     string    `json:"owner"`
	IsActive  *bool     `json:"is_active"`
}

func (a *Application) New(r *ApplicationRequest) {
	a.Name = r.Name
	a.Logo = r.Logo
	a.Owner = r.Owner
	a.IsActive = r.IsActive
}

func (a *Application) Validate() error {
	return nil
}

func (r *ApplicationUpdateRequest) NewUpdate() Map {
	return map[string]interface{}{}
}
