package model

import (
	utilitas "github.com/andrepriyanto10/favaa_mitra/utils"
	"gorm.io/gorm"
)

type Customers struct {
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
func (c *Customers) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = utilitas.NewUUIDString()

	return
}

type Address struct {
	ID          string `validate:"required,alphanumunicode"`
	CustomerID  string
	City        string `validate:"required"`
	Province    string `validate:"required"`
	Street      string `validate:"required"`
	County      string `validate:"required"`
	FullAddress string `validate:"required"`
}

func (c *Address) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = utilitas.NewUUIDString()

	return
}
