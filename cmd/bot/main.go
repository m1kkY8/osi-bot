package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/m1kkY8/osi-bot/pkg/api/htb"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/intents"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

var (
	users   []models.User
	perPage = 10
	pageMap = make(map[string]int)
)

func main() {
	godotenv.Load()

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content != "!osi" || m.Author.Bot {
			return
		}
		var err error
		users, err = htb.FetchUsers()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error fetching leaderboard")
			return
		}

		pageMap[m.Author.ID] = 0
		embed, buttons := embeds.LeaderboardEmbed(0, users)
		_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed:      embed,
			Components: buttons,
		})
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "error sending a message")
			return
		}
	})

	// Component (button) interaction handler
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}

		userID := i.Member.User.ID
		page := pageMap[userID]

		switch i.MessageComponentData().CustomID {
		case "leaderboard_next":
			page++
			if page*perPage >= len(users) {
				page = 0
			}
		case "leaderboard_prev":
			if page == 0 {
				page = (len(users) - 1) / perPage
			} else {
				page--
			}
		default:
			return
		}

		pageMap[userID] = page

		embed, components := embeds.LeaderboardEmbed(page, users)
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
			},
		})
		if err != nil {
			fmt.Println("Error updating message:", err)
		}
	})

	dg.Identify.Intents = intents.SetIntents()

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
