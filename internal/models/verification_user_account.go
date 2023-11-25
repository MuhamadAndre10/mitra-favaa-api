package models

import "time"

// VerificationData model
type VerificationData struct {
	Email     string    `gorm:"primaryKey;column:email;not null;unique"`
	Code      string    `gorm:"column:code;not null"`
	ExpiredAt time.Time `gorm:"column:expired_at;not null"`
}
