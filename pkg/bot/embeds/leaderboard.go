package embeds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func LeaderboardEmbed(page int, users []models.User) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	perPage := 10
	start := page * perPage
	end := min(start+perPage, len(users))

	embed := &discordgo.MessageEmbed{
		Title:  "OSI Leaderboard",
		Color:  0x2e2e2e,
		Fields: []*discordgo.MessageEmbedField{},
		Footer: &discordgo.MessageEmbedFooter{Text: "Use ⬅️ and ➡️ to scroll."},
	}

	for i, user := range users[start:end] {
		index := start + i + 1
		var medal string
		switch {
		case index == 1:
			medal = "🥇"
		case index == 2:
			medal = "🥈"
		case index == 3:
			medal = "🥉"
		default:
			medal = fmt.Sprintf("%2d.", index)
		}
		user.Name = fmt.Sprintf("%s %s", medal, user.Name)

		field := &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s | Points: %d", user.Name, user.Points),
			Value:  fmt.Sprintf("Rank: %s | User Owns: %d | Root Owns: %d", user.RankText, user.UserOwns, user.RootOwns),
			Inline: false,
		}
		embed.Fields = append(embed.Fields, field)
	}

	buttons := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Prev",
					Style:    discordgo.PrimaryButton,
					CustomID: "button_prev",
				},
				discordgo.Button{
					Label:    "Next",
					Style:    discordgo.PrimaryButton,
					CustomID: "button_next",
				},
			},
		},
	}

	return embed, buttons
}
