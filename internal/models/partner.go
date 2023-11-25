package models

import (
	utilitas "github.com/andrepriyanto10/favaa_mitra/internal/helpers"
	"gorm.io/gorm"
)

type Partner struct {
	ID           string  `gorm:"primaryKey"`
	FirstName    string  `gorm:"column:first_name"`
	LastName     string  `gorm:"column:last_name"`
	CustomerCode string  `gorm:"unique;column:customer_code"`
	Address      Address `gorm:"foreignKey:CustomerID;references:ID"`
	UserID       string
	CreatedAt    string `gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt    string `gorm:"column:updated_at;not null;autoUpdateTime"`
	DeletedAt    string `gorm:"column:deleted_at;not null"`
}

// BeforeCreate make a hook before insert into database
func (c *Partner) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = utilitas.NewUUIDString()

	return
}
