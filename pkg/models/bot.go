package models

import "github.com/bwmarrin/discordgo"

type Client struct {
	Users          []User
	BookstackUsers []BookstackUser
	DiscordSession *discordgo.Session
}

func NewClient(users []User, session *discordgo.Session) *Client {
	return &Client{
		Users:          []User{},
		DiscordSession: session,
	}
}
