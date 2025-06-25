package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/htb"
	"github.com/m1kkY8/osi-bot/pkg/models"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func teamGetRequestsSlashHandler(client *models.Client) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {
			util.RespondEphemeral(s, i.Interaction, "❌ You do not have permission to use this command.")
			return
		}
		joinReqs, err := htb.GetJoinRequests(client.GetTeamID())
		if err != nil {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("Failed to retrieve join requests: %v", err))
			return
		}
		if len(joinReqs) == 0 {
			util.RespondEphemeral(s, i.Interaction, "No requests found")
			return
		}
		content := fmt.Sprintf("Total join requests: %d\n", len(joinReqs))
		for _, req := range joinReqs {
			content += fmt.Sprintf("- Username: %s, User ID: %d, Request ID: %d\n", req.User.Name, req.User.ID, req.ID)
		}
		util.RespondEphemeral(s, i.Interaction, content)
	}
}
