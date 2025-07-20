package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/types"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

// Handler factory: returns a handler func for slash command
func bookUserSlashCommandHandler(client *types.Client, pages *types.Page) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {

			util.RespondEphemeral(s, i.Interaction, "❌ You do not have permission to use this command.")
			return
		}
		client.BookstackUsers = endpoints.BookApiListUsers()

		userID := i.Member.User.ID
		pages.PageMap[userID] = 0
		page := 0

		embed, components := embeds.ListUsersEmbed(page, client.BookstackUsers)

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
			},
		})
		if err != nil {
			fmt.Println("[ERROR] sending users embed:", err)
		}
	}
}
