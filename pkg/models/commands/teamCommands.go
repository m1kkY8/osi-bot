package commands

import (
	"github.com/bwmarrin/discordgo"
	commandFactory "github.com/m1kkY8/osi-bot/pkg/factories/commands"
)

func GetTeamCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "team",
			Description: "Manage HTB Team",
			Options:     commandFactory.GetAllHTBTeamCommands(),
		},
	}
}
