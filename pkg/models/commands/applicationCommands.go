package commands

import (
	"github.com/bwmarrin/discordgo"
)

// SetApplicationCommands initializes the application commands for the Discord bot client.
func GetApplicationCommands() []*discordgo.ApplicationCommand {
	alexandriaCommands := GetAlexandriaCommands()
	teamCommands := GetTeamCommands()

	return append(alexandriaCommands, teamCommands...)
}
