package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/htb"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func LeaderboardHandler(client *models.Client, pages *models.Page) {
	client.DiscordSession.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content != "!osi" || m.Author.Bot {
			return
		}
		var err error

		client.Users, err = htb.FetchUsers()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error fetching leaderboard")
			return
		}

		pages.PageMap[m.Author.ID] = 0
		embed, buttons := embeds.LeaderboardEmbed(0, client.Users)
		_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed:      embed,
			Components: buttons,
		})
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "error sending a message")
			return
		}
	})
}
