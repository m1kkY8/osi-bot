package models

import (
	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/bot/intents"
	commands "github.com/m1kkY8/osi-bot/pkg/models/commands"
	"github.com/m1kkY8/osi-bot/pkg/types"
)

// Re-export Client type for backward compatibility
type Client = types.Client

// NewClient creates a new Discord bot client
func NewClient(teamMembers []TeamMember, discordSession *discordgo.Session) *Client {
	return types.NewClient(teamMembers, discordSession)
}

// InitializeClient sets up the client with intents and commands
func InitializeClient(client *Client) {
	client.DiscordSession.Identify.Intents = intents.SetIntents()
	client.ApplicationCommands = commands.SetApplicationCommands()
	// Call the base Initialize method for environment variables
	client.Initialize()
}
