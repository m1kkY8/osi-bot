package factories

import (
	"github.com/m1kkY8/osi-bot/pkg/types"
)

// NewBookstackUser creates a new BookStack user with default values
func CreateBookstackUser(name, email, password string) *types.BookstackUser {
	return &types.BookstackUser{
		Name:       name,
		Email:      email,
		Password:   password,
		Roles:      []int{types.VIEWER},
		Language:   "en",
		SendInvite: false,
	}
}

// PromoteBookstackUserToEditor promotes a user to editor role
func PromoteBookstackUserToEditor(existingUser *types.BookstackUser) *types.BookstackUser {
	user := *existingUser // Copy to avoid mutating original
	user.Roles = []int{types.EDITOR}
	return &user
}
