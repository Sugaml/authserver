package domain

type Resource struct {
	BaseModel
	Enabled       bool
	Name          string
	DisplayName   string
	Description   string
	ApiSecrets    []ApiSecret   `gorm:"foreignkey:ApiResourceID"`
	ApiScopes     []ApiScope    `gorm:"foreignkey:ApiResourceID"`
	ApiProperties []ApiProperty `gorm:"foreignkey:ApiResourceID"`
}
