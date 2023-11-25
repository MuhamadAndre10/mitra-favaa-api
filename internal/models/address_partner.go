package models

import (
	utilitas "github.com/andrepriyanto10/favaa_mitra/internal/helpers"
	"gorm.io/gorm"
)

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
