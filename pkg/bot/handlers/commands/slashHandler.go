package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/types"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

type CommandHandler func(*discordgo.Session, *discordgo.InteractionCreate)

func HandleSlashCommand(
	client *types.Client,
	lbPages, bookstackPages *types.Page,
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	data := i.ApplicationCommandData()

	switch data.Name {
	case "team":
		handleTeamCommands(client, lbPages, s, i)
	case "alexandria":
		handleAlexandriaCommands(client, bookstackPages, s, i)
	default:
		fmt.Printf("[LOG] Unknown command: %s by %s\n", data.Name, i.Member.User.Username)
	}
}

func handleTeamCommands(client *types.Client, lbPages *types.Page, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if len(i.ApplicationCommandData().Options) == 0 {
		util.RespondEphemeral(s, i.Interaction, "Missing subcommand")
		return
	}

	sub := i.ApplicationCommandData().Options[0]
	fmt.Printf("[LOG] /team %s called by %s\n", sub.Name, i.Member.User.Username)

	// Map of command names to their handlers
	teamHandlers := map[string]CommandHandler{
		"invitations": teamGetRequestsSlashHandler(client),
		"accept":      teamAcceptSlashHandler(client),
		"reject":      teamRejectSlashHandler(client),
		"kick":        teamKickSlashHandler(client),
		"leaderboard": teamLeaderboardSlashHandler(client, lbPages),
	}

	if handler, exists := teamHandlers[sub.Name]; exists {
		handler(s, i)
	} else {
		util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("Unknown subcommand: %s", sub.Name))
	}
}

func handleAlexandriaCommands(client *types.Client, bookstackPages *types.Page, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if len(i.ApplicationCommandData().Options) == 0 {
		util.RespondEphemeral(s, i.Interaction, "Missing subcommand")
		return
	}

	sub := i.ApplicationCommandData().Options[0]
	fmt.Printf("[LOG] /alexandria %s called by %s\n", sub.Name, i.Member.User.Username)

	// Map of command names to their handlers
	alexandriaHandlers := map[string]CommandHandler{
		"register": registerUserSlashHandler(client),
		"remove":   deleteUserSlashHandler(client),
		"users":    bookUserSlashCommandHandler(client, bookstackPages),
		"update":   updateUserSlashHandler(client),
	}

	if handler, exists := alexandriaHandlers[sub.Name]; exists {
		handler(s, i)
	} else {
		util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("Unknown subcommand: %s", sub.Name))
	}
}
