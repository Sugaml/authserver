package domain

import "time"

type Resource struct {
	BaseModel
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type ResourceRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type ResourceListRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type ResourceUpdateRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type ResourceResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Description string    `json:"description"`
	Enabled     bool      `json:"enabled"`
}

func (a *Resource) New(r *ResourceRequest) {
	a.Name = r.Name
	a.DisplayName = r.DisplayName
	a.Description = r.Description
	a.Enabled = r.Enabled
}

func (a *Resource) Validate() error {
	return nil
}

func (a *ResourceUpdateRequest) NewUpdate() Map {
	m := Map{}
	m["name"] = a.Name
	m["display_name"] = a.DisplayName
	m["description"] = a.Description
	m["enabled"] = a.Enabled
	return m
}
