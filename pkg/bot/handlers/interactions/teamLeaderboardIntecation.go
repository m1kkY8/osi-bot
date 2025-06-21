package interactions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func LeaderboardInteraction(client *models.Client, pages *models.Page) {
	client.DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}
		fmt.Println(i.MessageComponentData().CustomID)

		userID := i.Member.User.ID
		page := pages.PageMap[userID]

		switch i.MessageComponentData().CustomID {
		case "button_next":
			page++
			if page*pages.PerPage >= len(client.TeamMembers) {
				page = 0
			}
		case "button_prev":
			if page == 0 {
				page = (len(client.TeamMembers) - 1) / pages.PerPage
			} else {
				page--
			}
		default:
			return
		}

		pages.PageMap[userID] = page

		embed, components := embeds.LeaderboardEmbed(page, client.TeamMembers)
		currentPage := page + 1
		footerText := fmt.Sprintf("Page %d/%d", currentPage, (len(client.TeamMembers)+pages.PerPage-1)/pages.PerPage)

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
