package utils

import "fmt"

// CustomError represents a custom error type
type CustomError struct {
	Code    int
	Message string
}

// Implement the error interface for CustomError
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
