package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

// Handler factory: returns a handler func for slash command
func BookUserSlashCommandHandler(client *models.Client, pages *models.Page) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand || i.ApplicationCommandData().Name != "users" {
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
