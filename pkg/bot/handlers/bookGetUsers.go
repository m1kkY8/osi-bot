package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func GetUsers(client *models.Client, pages *models.Page) {
	client.DiscordSession.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot || !strings.HasPrefix(m.Content, "!bookuser") {
			return
		}

		client.BookstackUsers = endpoints.BookApiListUsers()

		// Initialize the page number to 0 for this user
		userID := m.Author.ID
		pages.PageMap[userID] = 0
		page := 0

		embed, components := embeds.ListUsersEmbed(page, client.BookstackUsers)

		_, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed:      embed,
			Components: components,
		})
		if err != nil {
			fmt.Println("[ERROR] sending users embed:", err)
		}
	})
}
