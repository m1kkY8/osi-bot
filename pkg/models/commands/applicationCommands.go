package models

import (
	"github.com/bwmarrin/discordgo"
)

// SetApplicationCommands initializes the application commands for the Discord bot client.
func SetApplicationCommands() []*discordgo.ApplicationCommand {
	alexandriaCommands := getAlexandriaCommands()
	teamCommands := GetTeamCommands()

	return append(alexandriaCommands, teamCommands...)
}
