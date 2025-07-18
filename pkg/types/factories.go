package types

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/gubarz/gohtb"
)

// NewBookstackUser creates a new BookStack user with default values
func NewBookstackUser(name, email, password string) *BookstackUser {
	return &BookstackUser{
		Name:       name,
		Email:      email,
		Password:   password,
		Roles:      []int{VIEWER},
		Language:   "en",
		SendInvite: false,
	}
}

// NewClient creates a new Discord bot client
func NewClient(teamMembers []TeamMember, discordSession *discordgo.Session) *Client {
	return &Client{
		Context:             context.Background(),
		HTBClient:           gohtb.Client{},
		TeamMembers:         teamMembers,
		BookstackUsers:      []BookstackUser{},
		DiscordSession:      discordSession,
		ApplicationCommands: []*discordgo.ApplicationCommand{},
		Intents:             []discordgo.Intent{},
		GuildID:             "",
		AdminRoleID:         "",
		TeamID:              "",
	}
}

// NewPage creates a new pagination object
func NewPage(currentPage, perPage, totalPages int, pageMap map[string]int) *Page {
	return &Page{
		CurrentPage: currentPage,
		PerPage:     perPage,
		TotalPages:  totalPages,
		PageMap:     pageMap,
	}
}

// PromoteBookstackUserToEditor promotes a user to editor role
func PromoteBookstackUserToEditor(existingUser *BookstackUser) *BookstackUser {
	user := *existingUser // Copy to avoid mutating original
	user.Roles = []int{EDITOR}
	return &user
}
