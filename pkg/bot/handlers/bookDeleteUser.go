package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func DeleteUser(client *models.Client) {
	client.DiscordSession.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot || !strings.HasPrefix(m.Content, "!removeuser") {
			return
		}

		// Split message into command and args
		args := strings.Fields(m.Content)
		if len(args) < 2 {
			s.ChannelMessageSend(m.ChannelID, "❌ Usage: `!removeuser <user_id>`")
			return
		}

		// Convert argument to integer
		id, err := strconv.Atoi(args[1])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "❌ Invalid ID. Please enter a valid number.")
			return
		}

		// Call your deletion logic
		endpoints.BookApiDeleteUser(id)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("✅ Attempting to delete user with ID: %d", id))
	})
}
