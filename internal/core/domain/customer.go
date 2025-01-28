package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Username                  string
	PasswordHash              string
	SecurityStamp             string
	ConcurrencyStamp          string
	Email                     string
	Domain                    string `json:"domain"`
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

type CustomerRequest struct {
	Username             string `json:"username"`
	Domain               string `json:"domain"`
	PasswordHash         string `json:"password"`
	SecurityStamp        string `json:"security_stamp"`
	Email                string `json:"email"`
	EmailConfirmed       bool   `json:"email_confirmed"`
	PhoneNumber          string `json:"phone_number"`
	PhoneNumberConfirmed bool   `json:"phone_number_confirmed"`
	TwoFactorEnabled     bool   `json:"two_factor_enable"`
}

type CustomerUpdateRequest struct {
	Username             string `json:"username"`
	PasswordHash         string `json:"password"`
	Domain               string `json:"domain"`
	SecurityStamp        string `json:"security_stamp"`
	Email                string `json:"email"`
	EmailConfirmed       bool   `json:"email_confirmed"`
	PhoneNumber          string `json:"phone_number"`
	PhoneNumberConfirmed *bool  `json:"phone_number_confirmed"`
	TwoFactorEnabled     *bool  `json:"two_factor_enable"`
}

type ListCustomerRequest struct {
	ListRequest
	Username             string `form:"username"`
	PasswordHash         string `form:"password"`
	SecurityStamp        string `form:"security_stamp"`
	Email                string `form:"email"`
	EmailConfirmed       bool   `form:"email_confirmed"`
	PhoneNumber          string `form:"phone_number"`
	PhoneNumberConfirmed bool   `form:"phone_number_confirmed"`
	TwoFactorEnabled     bool   `form:"two_factor_enable"`
}

func (a *Customer) New(r *CustomerRequest) {
	a.Username = r.Username
	a.Email = r.Email
	a.Domain = r.Domain
	a.EmailConfirmed = r.EmailConfirmed
	a.PasswordHash = r.PasswordHash
	a.PhoneNumber = r.PhoneNumber
	a.PhoneNumberConfirmed = r.PhoneNumberConfirmed
	a.TwoFactorEnabled = r.TwoFactorEnabled
}

func (a *Customer) Validate() error {
	return nil
}

func (r *CustomerUpdateRequest) NewUpdate() Map {
	return map[string]interface{}{}
}

type CustomerResponse struct {
	ID                   string    `json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	Username             string    `json:"username"`
	PasswordHash         string    `json:"password"`
	SecurityStamp        string    `json:"security_stamp"`
	Email                string    `json:"email"`
	EmailConfirmed       bool      `json:"email_confirmed"`
	PhoneNumber          string    `json:"phone_number"`
	PhoneNumberConfirmed bool      `json:"phone_number_confirmed"`
	TwoFactorEnabled     bool      `json:"two_factor_enable"`
}
