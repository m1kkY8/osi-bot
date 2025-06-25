package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/htb"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/models"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func teamLeaderboardSlashHandler(client *models.Client, pages *models.Page) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var err error

		client.TeamMembers, err = htb.FetchUsers(client.GetTeamID())
		if err != nil {
			_ = util.RespondEphemeral(s, i.Interaction, "Error fetching leaderboard")
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
