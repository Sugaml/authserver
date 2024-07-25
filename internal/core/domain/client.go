package domain

type Client struct {
	ID                 string `gorm:"primary_key"`
	Enabled            bool
	ClientID           string
	ProtocolType       string
	ClientName         string
	Description        string
	ClientUri          string
	LogoutUri          string
	EnabledLocalLogin  bool
	ClientCorsOrigins  []ClientCorsOrigin  `gorm:"foreignkey:ClientID"`
	ClientGrantTypes   []ClientGrantType   `gorm:"foreignkey:ClientID"`
	ClientProperties   []ClientProperty    `gorm:"foreignkey:ClientID"`
	ClientRedirectUris []ClientRedirectUri `gorm:"foreignkey:ClientID"`
	ClientScopes       []ClientScope       `gorm:"foreignkey:ClientID"`
	ClientSecrets      []ClientSecret      `gorm:"foreignkey:ClientID"`
}
