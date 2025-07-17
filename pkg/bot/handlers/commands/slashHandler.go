package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/models"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func HandleSlashCommand(
	client *models.Client,
	lbPages, bookstackPages *models.Page,
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

func handleTeamCommands(client *models.Client, lbPages *models.Page, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if len(i.ApplicationCommandData().Options) == 0 {
		util.RespondEphemeral(s, i.Interaction, "Missing subcommand")
		return
	}

	sub := i.ApplicationCommandData().Options[0]
	fmt.Printf("[LOG] /team %s called by %s\n", sub.Name, i.Member.User.Username)

	var handler func(*discordgo.Session, *discordgo.InteractionCreate)
	switch sub.Name {
	case "invitations":
		handler = teamGetRequestsSlashHandler(client)
	case "accept":
		handler = teamAcceptSlashHandler(client)
	case "reject":
		handler = teamRejectSlashHandler(client)
	case "kick":
		handler = teamKickSlashHandler(client)
	case "leaderboard":
		handler = teamLeaderboardSlashHandler(client, lbPages)
	default:
		util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("Unknown subcommand: %s", sub.Name))
	}

	if handler != nil {
		handler(s, i)
	}
}

func handleAlexandriaCommands(client *models.Client, bookstackPages *models.Page, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if len(i.ApplicationCommandData().Options) == 0 {
		util.RespondEphemeral(s, i.Interaction, "Missing subcommand")
		return
	}

	sub := i.ApplicationCommandData().Options[0]
	fmt.Printf("[LOG] /alexandria %s called by %s\n", sub.Name, i.Member.User.Username)

	var handler func(*discordgo.Session, *discordgo.InteractionCreate)
	switch sub.Name {
	case "register":
		handler = registerUserSlashHandler(client)
	case "remove":
		handler = deleteUserSlashHandler(client)
	case "users":
		handler = bookUserSlashCommandHandler(client, bookstackPages)
	case "update":
		handler = updateUserSlashHandler(client)
	default:
		util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("Unknown subcommand: %s", sub.Name))
	}

	if handler != nil {
		handler(s, i)
	}
}
