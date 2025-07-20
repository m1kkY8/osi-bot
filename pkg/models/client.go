package models

import (
	"github.com/m1kkY8/osi-bot/pkg/bot/intents"
	"github.com/m1kkY8/osi-bot/pkg/models/commands"
	"github.com/m1kkY8/osi-bot/pkg/types"
)

// InitializeClient sets up the client with intents and commands
func InitializeClient(client *types.Client) {
	client.DiscordSession.Identify.Intents = intents.SetIntents()

	var applicationCommands = commands.GetApplicationCommands()
	client.SetApplicationCommands(applicationCommands)
	// Call the base Initialize method for environment variables
	client.Initialize()
}
