package embeds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/types"
)

func LeaderboardEmbed(page int, users []types.TeamMember) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	perPage := 10
	start := page * perPage
	end := min(start+perPage, len(users))

	currentPage := page + 1
	footerText := fmt.Sprintf("Page %d/%d", currentPage, (len(users)+perPage-1)/perPage)

	embed := &discordgo.MessageEmbed{
		Title:  "OSI Leaderboard",
		Color:  0x2e2e2e,
		Fields: []*discordgo.MessageEmbedField{},
		Footer: &discordgo.MessageEmbedFooter{Text: footerText},
	}

	for i, user := range users[start:end] {
		index := start + i + 1
		var medal string
		switch index {
		case 1:
			medal = "🥇"
		case 2:
			medal = "🥈"
		case 3:
			medal = "🥉"
		default:
			medal = fmt.Sprintf("%2d.", index)
		}
		user.Name = fmt.Sprintf("%s %s", medal, user.Name)

		field := &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s | Points: %d | ID: %d", user.Name, user.Points, user.ID),
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
