package auth

import "github.com/google/uuid"

func GeneratePassword() string {
	return uuid.NewString()
}
