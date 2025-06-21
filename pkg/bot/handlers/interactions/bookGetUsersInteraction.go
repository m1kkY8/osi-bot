package interactions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func UserListInteraction(client *models.Client, pages *models.Page) {
	client.DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}

		customID := i.MessageComponentData().CustomID
		if customID != "button_next_book" && customID != "button_prev_book" {
			fmt.Println("unknown button")
			return
		}

		userID := i.Member.User.ID
		page := pages.PageMap[userID]

		switch customID {
		case "button_next_book":
			page++
			if page*pages.PerPage >= len(client.BookstackUsers) {
				page = 0
			}
		case "button_prev_book":
			if page == 0 {
				page = (len(client.BookstackUsers) - 1) / pages.PerPage
			} else {
				page--
			}
		}

		pages.PageMap[userID] = page

		embed, components := embeds.ListUsersEmbed(page, client.BookstackUsers)
		currentPage := page + 1
		footerText := fmt.Sprintf("Page %d/%d", currentPage, (len(client.BookstackUsers)+pages.PerPage-1)/pages.PerPage)

		embed.Footer = &discordgo.MessageEmbedFooter{Text: footerText}
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
}
