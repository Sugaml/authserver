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
