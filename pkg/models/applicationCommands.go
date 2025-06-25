package models

import (
	"github.com/bwmarrin/discordgo"
)

// SetApplicationCommands initializes the application commands for the Discord bot client.
func SetApplicationCommands() []*discordgo.ApplicationCommand {
	alexandriaCommands := getAlexandriaCommands()
	teamCommands := getTeamCommands()

	return append(alexandriaCommands, teamCommands...)
}
