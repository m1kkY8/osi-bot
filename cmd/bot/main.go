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
		msg, _ := s.ChannelMessageSendEmbed(m.ChannelID, embeds.LeaderboardEmbed(0, users))

		s.MessageReactionAdd(m.ChannelID, msg.ID, "⬅️")
		s.MessageReactionAdd(m.ChannelID, msg.ID, "➡️")
	})

	dg.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
		if r.UserID == s.State.User.ID || r.Emoji.Name != "⬅️" && r.Emoji.Name != "➡️" {
			return
		}

		page, exists := pageMap[r.UserID]
		if !exists {
			return
		}

		if r.Emoji.Name == "➡️" {
			page++
			if page*perPage >= len(users) {
				page = 0 // Loop back to the start
			}
		} else if r.Emoji.Name == "⬅️" {
			if page == 0 {
				page = (len(users) / perPage) // Go to the last page
			} else {
				page--
			}
		}

		pageMap[r.UserID] = page
		s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
		s.ChannelMessageEditEmbed(r.ChannelID, r.MessageID, embeds.LeaderboardEmbed(page, users))
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
