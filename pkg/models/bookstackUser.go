package models

import "github.com/m1kkY8/osi-bot/pkg/types"

// Re-export types for backward compatibility
type BookstackUser = types.BookstackUser
type BookstackUserResponse = types.BookstackUserResponse

// CreateBookstackUser creates a new BookStack user with default values
func CreateBookstackUser(name, email, password string) *BookstackUser {
	return types.NewBookstackUser(name, email, password)
}
