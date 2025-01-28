package domain

import (
	"time"

	"github.com/google/uuid"
)

type Map map[string]interface{}

type Client struct {
	BaseModel
	Enabled            bool
	ClientID           string
	ProtocolType       string
	ClientName         string
	Description        string
	ClientUri          string
	LogoutUri          string
	ApplicationID      string
	EnabledLocalLogin  bool
	ClientCorsOrigins  []ClientCorsOrigin  `gorm:"foreignkey:ClientID"`
	ClientGrantTypes   []ClientGrantType   `gorm:"foreignkey:ClientID"`
	ClientProperties   []ClientProperty    `gorm:"foreignkey:ClientID"`
	ClientRedirectUris []ClientRedirectUri `gorm:"foreignkey:ClientID"`
	ClientScopes       []ClientScope       `gorm:"foreignkey:ClientID"`
	ClientSecrets      []ClientSecret      `gorm:"foreignkey:ClientID"`
}

type ClientRequest struct {
	ID                string `json:"id"`
	Enabled           bool
	ApplicationID     string
	ClientID          string `json:"client_id"`
	ProtocolType      string
	ClientName        string `json:"client_name"`
	Description       string
	ClientUri         string
	LogoutUri         string
	EnabledLocalLogin bool
}

type ClientUpdateRequest struct {
	ID                string `json:"id"`
	Enabled           bool
	ApplicationID     string
	ClientID          string
	ProtocolType      string
	ClientName        string
	Description       string
	ClientUri         string
	LogoutUri         string
	EnabledLocalLogin bool
}

func (a *Client) New(r *ClientRequest) {

	a.ApplicationID = r.ApplicationID
	a.ClientID = r.ClientID
	if r.ClientID == "" {
		a.ClientID = uuid.New().String()[:16]
	}
}

func (a *Client) Validate() error {
	return nil
}

func (r *ClientUpdateRequest) NewUpdate() Map {
	return map[string]interface{}{}
}

type ListRequest struct {
	Page          int64  `json:"page"`
	Size          int64  `json:"size"`
	SortColumn    string `json:"sort_column"`
	SortDirection string `json:"sort_direction"`
	Query         string `json:"query"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}

type ClientListRequest struct {
	ListRequest
	Enabled           bool
	ClientID          string
	ProtocolType      string
	ApplicationID     string
	ClientName        string
	Description       string
	ClientUri         string
	LogoutUri         string
	EnabledLocalLogin bool
}

type ClientResponse struct {
	ID                string    `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	Enabled           bool
	ApplicationID     string
	ClientID          string
	ProtocolType      string
	ClientName        string
	Description       string
	ClientUri         string
	LogoutUri         string
	EnabledLocalLogin bool
}
