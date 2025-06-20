package models

type BookstackUser struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Roles      []int  `json:"roles"`
	Language   string `json:"language"`
	SendInvite bool   `json:"sendInvite"`
}
