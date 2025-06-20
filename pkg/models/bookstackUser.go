package models

import "github.com/m1kkY8/osi-bot/pkg/models/roles"

type BookstackUser struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Roles      []int  `json:"roles"`
	Language   string `json:"language"`
	SendInvite bool   `json:"sendInvite"`
}

type BookstackUserResponse struct {
	Data  []BookstackUser `json:"data"`
	Total int             `json:"total"`
}

func CreateBookstackUser(name, email, password string) *BookstackUser {
	user := &BookstackUser{
		Name:       name,
		Email:      email,
		Password:   password,
		Roles:      []int{roles.VIEWER},
		Language:   "en",
		SendInvite: false,
	}
	return user
}
