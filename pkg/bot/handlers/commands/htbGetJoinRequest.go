package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/htb"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func TeamGetRequestsSlashHandler(client *models.Client) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "❌ You do not have permission to use this command.",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}
		joinReqs, err := htb.GetJoinRequests()
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Failed to retrieve join requests: %v", err),
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}
		if len(joinReqs) == 0 {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "No join requests found.",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}
		content := fmt.Sprintf("Total join requests: %d\n", len(joinReqs))
		for _, req := range joinReqs {
			content += fmt.Sprintf("- Username: %s, User ID: %d, Request ID: %d\n", req.User.Name, req.User.ID, req.ID)
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}
