package commands

import (
	"github.com/bwmarrin/discordgo"
	commandFactory "github.com/m1kkY8/osi-bot/pkg/factories/commands"
)

func GetAlexandriaCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "alexandria",
			Description: "Manage Alexandria book stack",
			Options:     commandFactory.GetAllAlexandriaCommands(),
		},
	}
}
