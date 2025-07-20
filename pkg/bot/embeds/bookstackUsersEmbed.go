package embeds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/types"
)

func ListUsersEmbed(page int, users []types.BookstackUser) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	perPage := 10
	start := page * perPage
	end := min(start+perPage, len(users))

	embed := &discordgo.MessageEmbed{
		Title:  "Alexandria Users",
		Color:  0x5865f2,
		Fields: []*discordgo.MessageEmbedField{},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Page %d/%d", page+1, (len(users)+perPage-1)/perPage),
		},
	}

	for _, user := range users[start:end] {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("Username: %s (ID: %d)", user.Name, user.ID),
			Value:  fmt.Sprintf("Email: %s", user.Email),
			Inline: false,
		})
	}

	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Prev",
					Style:    discordgo.PrimaryButton,
					CustomID: "button_prev_book",
				},
				discordgo.Button{
					Label:    "Next",
					Style:    discordgo.PrimaryButton,
					CustomID: "button_next_book",
				},
			},
		},
	}
	return embed, components
}
