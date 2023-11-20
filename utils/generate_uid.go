package utils

import "github.com/google/uuid"

func NewUUIDString() string {
	uid := uuid.New().String()
	return uid
}
