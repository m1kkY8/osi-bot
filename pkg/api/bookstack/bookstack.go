package bookstack

import (
	"github.com/m1kkY8/osi-bot/pkg/models"
	"github.com/m1kkY8/osi-bot/pkg/models/roles"
)

func CreateBookstackUser(name, email, password string) *models.BookstackUser {
	user := &models.BookstackUser{
		Name:       name,
		Email:      email,
		Password:   password,
		Roles:      []int{roles.VIEWER},
		Language:   "en",
		SendInvite: false,
	}
	return user
}
