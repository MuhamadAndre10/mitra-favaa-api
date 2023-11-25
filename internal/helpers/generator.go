package helpers

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func CodeVerification() string {
	rand.NewSource(time.Now().UnixNano())

	// Generate a random code consisting of 5 numbers
	code := ""
	for i := 0; i < 5; i++ {
		// Append a random digit to the code
		code += fmt.Sprintf("%d", rand.Intn(10))
	}

	return code
}

func NewUUIDString() string {
	uid := uuid.New().String()
	return uid
}
