package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/htb"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func LeaderboardSlashHandler(client *models.Client, pages *models.Page) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand || i.ApplicationCommandData().Name != "leaderboard" {
			return
		}
		var err error

		client.TeamMembers, err = htb.FetchUsers()
		if err != nil {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error fetching leaderboard",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}

		userID := i.Member.User.ID
		pages.PageMap[userID] = 0
		embed, buttons := embeds.LeaderboardEmbed(0, client.TeamMembers)
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: buttons,
			},
		})
		if err != nil {
			fmt.Println("error sending a message:", err)
			return
		}
	}
}
