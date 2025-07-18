package types

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/gubarz/gohtb"
)

// Client represents the main bot client with all necessary components
type Client struct {
	Context             context.Context
	HTBClient           gohtb.Client
	TeamMembers         []TeamMember
	BookstackUsers      []BookstackUser
	DiscordSession      *discordgo.Session
	ApplicationCommands []*discordgo.ApplicationCommand
	Intents             []discordgo.Intent
	GuildID             string
	AdminRoleID         string
	TeamID              string
}

// Page represents pagination information for interactive components
type Page struct {
	CurrentPage int
	PerPage     int
	TotalPages  int
	PageMap     map[string]int
}
