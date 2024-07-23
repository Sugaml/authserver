package domain

import (
	"time"
)

type RoleClaim struct {
	ID         string `gorm:"primary_key"`
	RoleID     string
	ClaimType  string
	ClaimValue string
}

type Role struct {
	ID               string `gorm:"primary_key"`
	Name             string
	NormalizedName   string
	ConcurrencyStamp string
	RoleClaims       []RoleClaim       `gorm:"foreignkey:RoleID"`
	RoleBundlesRoles []RoleBundlesRole `gorm:"foreignkey:RoleID"`
}

type UserClaim struct {
	ID         string `gorm:"primary_key"`
	UserID     string
	ClaimType  string
	ClaimValue string
}

type UserLogin struct {
	LoginProvider       string
	ProviderKey         string
	ProviderDisplayName string
	UserID              string
}

type UserRole struct {
	UserID string
	RoleID string
}

type UserToken struct {
	UserID        string
	LoginProvider string
	Name          string
	Value         string
}

type User struct {
	ID                   uint        `gorm:"primary_key" json:"id"`
	UserName             string      `json:"user_name"`
	NormalizedUserName   string      `json:"normalized_user_name"`
	Email                string      `json:"email"`
	EmailConfirmed       bool        `json:"email_confirmed"`
	Password             string      `json:"password_hash"`
	SecurityStamp        string      `json:"security_stamp"`
	PhoneNumber          string      `json:"phone_number"`
	PhoneNumberConfirmed bool        `json:"phone_number_confirmed"`
	TwoFactorEnabled     bool        `json:"two_factor_enabled"`
	LockoutEnd           *time.Time  `json:"lockout_end"`
	LockoutEnabled       bool        `json:"lockout_enabled"`
	AccessFailedCount    int         `json:"access_failed_count"`
	ConcurrencyStamp     string      `json:"concurrency_stamp"`
	UserClaims           []UserClaim `gorm:"foreignkey:UserID" json:"user_claims"`
	UserLogins           []UserLogin `gorm:"foreignkey:UserID" json:"user_logins"`
	UserRoles            []UserRole  `gorm:"foreignkey:UserID" json:"user_roles"`
	UserTokens           []UserToken `gorm:"foreignkey:UserID" json:"user_tokens"`
}

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

type ClientCorsOrigin struct {
	ID       string `gorm:"primary_key"`
	Origin   string
	ClientID string
}

type ClientGrantType struct {
	ID        string `gorm:"primary_key"`
	GrantType string
	ClientID  string
}

type ClientProperty struct {
	ID       string `gorm:"primary_key"`
	Key      string
	Value    string
	ClientID string
}

type ClientRedirectUri struct {
	ID          string `gorm:"primary_key"`
	RedirectUri string
	ClientID    string
}

type ClientScope struct {
	ID       string `gorm:"primary_key"`
	Scope    string
	ClientID string
}

type ClientSecret struct {
	ID          string `gorm:"primary_key"`
	Description string
	Value       string
	Expiration  *time.Time
	ClientID    string
}

type Customer struct {
	ID                        string `gorm:"primary_key"`
	CreatedUtc                time.Time
	UpdatedUtc                time.Time
	Username                  string
	PasswordHash              string
	SecurityStamp             string
	ConcurrencyStamp          string
	Email                     string
	EmailConfirmed            bool
	PhoneNumber               string
	PhoneNumberConfirmed      bool
	TwoFactorEnabled          bool
	LockoutEnd                *time.Time
	LockoutEnabled            bool
	AccessFailedCount         int
	CustomerExternalDomains   []CustomerExternalDomain   `gorm:"foreignkey:CustomerID"`
	CustomerExternalGroups    []CustomerExternalGroup    `gorm:"foreignkey:CustomerID"`
	CustomerIdentityProviders []CustomerIdentityProvider `gorm:"foreignkey:CustomerID"`
	AuditMessages             []AuditMessage             `gorm:"foreignkey:CustomerID"`
}

type CustomerExternalDomain struct {
	ID         string `gorm:"primary_key"`
	Domain     string
	CustomerID string
}

type CustomerExternalGroup struct {
	ID         string `gorm:"primary_key"`
	Name       string
	Type       string
	CustomerID string
}

type CustomerIdentityProvider struct {
	ID              string `gorm:"primary_key"`
	CreatedUtc      time.Time
	UpdatedUtc      time.Time
	Key             string
	DisplayName     string
	MetadataAddress string
	ProtocolType    string
	CustomerID      string
}

type AuditMessage struct {
	ID         string `gorm:"primary_key"`
	CreatedUtc time.Time
	UpdatedUtc time.Time
	User       string
	Message    string
	CustomerID string
}

type DeviceCode struct {
	UserCode     string
	DeviceCode   string
	SubjectID    string
	ClientID     string
	CreationTime time.Time
	Expiration   time.Time
	Data         string
}

type Key struct {
	ID                string `gorm:"primary_key"`
	Version           int
	Created           time.Time
	Use               string
	Algorithm         string
	IsX509Certificate bool
	DataProtected     bool
	Data              string
}

type PersistedGrant struct {
	Key          string `gorm:"primary_key"`
	Type         string
	SubjectID    string
	ClientID     string
	CreationTime time.Time
	Expiration   time.Time
	Data         string
}

type ApiResource struct {
	ID            string `gorm:"primary_key"`
	Enabled       bool
	Name          string
	DisplayName   string
	Description   string
	ApiSecrets    []ApiSecret   `gorm:"foreignkey:ApiResourceID"`
	ApiScopes     []ApiScope    `gorm:"foreignkey:ApiResourceID"`
	ApiProperties []ApiProperty `gorm:"foreignkey:ApiResourceID"`
}

type ApiProperty struct {
	ID            string `gorm:"primary_key"`
	Key           string
	Value         string
	ApiResourceID string
}

type ApiScope struct {
	ID                      string `gorm:"primary_key"`
	Name                    string
	DisplayName             string
	Description             string
	Required                bool
	Emphasize               bool
	ShowInDiscoveryDocument bool
	ApiResourceID           string
	ApiScopeClaims          []ApiScopeClaim `gorm:"foreignkey:ApiScopeID"`
}

type ApiScopeClaim struct {
	ID         string `gorm:"primary_key"`
	Type       string
	ApiScopeID string
}

type ApiSecret struct {
	ID            string `gorm:"primary_key"`
	Description   string
	Value         string
	Expiration    *time.Time
	ApiResourceID string
}

type IdentityResource struct {
	ID                      string `gorm:"primary_key"`
	Enabled                 bool
	Name                    string
	DisplayName             string
	Description             string
	Required                bool
	Emphasize               bool
	ShowInDiscoveryDocument bool
	IdentityClaims          []IdentityClaim    `gorm:"foreignkey:IdentityResourceID"`
	IdentityProperties      []IdentityProperty `gorm:"foreignkey:IdentityResourceID"`
}

type IdentityClaim struct {
	ID                 string `gorm:"primary_key"`
	Type               string
	IdentityResourceID string
}

type IdentityProperty struct {
	ID                 string `gorm:"primary_key"`
	Key                string
	Value              string
	IdentityResourceID string
}

type RoleBundle struct {
	ID               string `gorm:"primary_key"`
	CreatedUtc       time.Time
	UpdatedUtc       time.Time
	Name             string
	Description      string
	RoleBundlesRoles []RoleBundlesRole `gorm:"foreignkey:RoleBundleID"`
}

type RoleBundlesRole struct {
	ID           string `gorm:"primary_key"`
	RoleBundleID string
	RoleID       string
}

type Tenant struct {
	ID                 string `gorm:"primary_key"`
	CreatedUtc         time.Time
	UpdatedUtc         time.Time
	Name               string
	Description        string
	RoleBundlesTenants []RoleBundlesTenant `gorm:"foreignkey:TenantID"`
}

type RoleBundlesTenant struct {
	ID           string `gorm:"primary_key"`
	RoleBundleID string
	TenantID     string
}

type Setting struct {
	ID              string `gorm:"primary_key"`
	CreatedUtc      time.Time
	UpdatedUtc      time.Time
	Name            string
	ValueStringUtf8 string
	GlobalDefault   bool
}

type EFMigrationHistory struct {
	MigrationId    string `gorm:"primary_key"`
	ProductVersion string
}

type DataProtection struct {
	ID         string `gorm:"primary_key"`
	CreatedUtc time.Time
	UpdatedUtc time.Time
	Purpose    string
	Data       string
}
