package models

import "github.com/bwmarrin/discordgo"

type Client struct {
	TeamMembers    []TeamMember
	BookstackUsers []BookstackUser
	DiscordSession *discordgo.Session
}

func NewClient(users []TeamMember, session *discordgo.Session) *Client {
	return &Client{
		TeamMembers:    []TeamMember{},
		DiscordSession: session,
	}
}
