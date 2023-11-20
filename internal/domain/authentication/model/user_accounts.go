package model

import (
	utilitas "github.com/andrepriyanto10/favaa_mitra/utils"
	"gorm.io/gorm"
	"time"
)

type UserAccounts struct {
	ID             string    `gorm:"primaryKey;<-create;column:id;not null"`
	Username       string    `gorm:"column:username;not null"`
	Email          string    `gorm:"column:email;not null;unique"`
	Avatar         string    `gorm:"column:avatar;not null"`
	Status         string    `gorm:"column:status;not null"`
	Phone          string    `gorm:"column:phone;not null"`
	DateRegistered time.Time `gorm:"column:date_registered"`
	Password       string    `gorm:"column:password;not null"`
	Customer       Customers `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;autoUpdateTime"`
	DeletedAt      time.Time `gorm:"column:deleted_at;not null"`
}

// BeforeCreate make a hook before insert into database
func (u *UserAccounts) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = utilitas.NewUUIDString()

	return
}
