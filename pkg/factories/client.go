package factories

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/gubarz/gohtb"
	"github.com/m1kkY8/osi-bot/pkg/types"
)

// NewClient creates a new Discord bot client
func NewClient(teamMembers []types.TeamMember, discordSession *discordgo.Session) *types.Client {
	return &types.Client{
		Context:             context.Background(),
		HTBClient:           gohtb.Client{},
		TeamMembers:         teamMembers,
		BookstackUsers:      []types.BookstackUser{},
		DiscordSession:      discordSession,
		ApplicationCommands: []*discordgo.ApplicationCommand{},
		Intents:             []discordgo.Intent{},
		GuildID:             "",
		AdminRoleID:         "",
		TeamID:              "",
	}
}

// NewPage creates a new pagination object
func NewPage(currentPage, perPage, totalPages int, pageMap map[string]int) *types.Page {
	return &types.Page{
		CurrentPage: currentPage,
		PerPage:     perPage,
		TotalPages:  totalPages,
		PageMap:     pageMap,
	}
}
