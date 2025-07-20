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
	case types.CommandTeam:
		handleTeamCommands(client, lbPages, s, i)
	case types.CommandAlexandria:
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
		types.SubcommandInvitations: teamGetRequestsSlashHandler(client),
		types.SubcommandAccept:      teamAcceptSlashHandler(client),
		types.SubcommandReject:      teamRejectSlashHandler(client),
		types.SubcommandKick:        teamKickSlashHandler(client),
		types.SubcommandLeaderboard: teamLeaderboardSlashHandler(client, lbPages),
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
		types.SubcommandRegister: registerUserSlashHandler(client),
		types.SubcommandRemove:   deleteUserSlashHandler(client),
		types.SubcommandUsers:    bookUserSlashCommandHandler(client, bookstackPages),
		types.SubcommandUpdate:   updateUserSlashHandler(client),
	}

	if handler, exists := alexandriaHandlers[sub.Name]; exists {
		handler(s, i)
	} else {
		util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("Unknown subcommand: %s", sub.Name))
	}
}
