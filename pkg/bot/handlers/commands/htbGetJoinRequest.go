package commands

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/types"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func teamGetRequestsSlashHandler(client *types.Client) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {
			util.RespondEphemeral(s, i.Interaction, "❌ You do not have permission to use this command.")
			return
		}
		teamID, _ := strconv.Atoi(client.GetTeamID())
		result, err := client.HTBClient.Teams.Team(teamID).Invitations(client.Context)
		if err != nil {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("Failed to retrieve join requests: %v", err))
			return
		}
		if len(result.Data) == 0 {
			util.RespondEphemeral(s, i.Interaction, "No requests found")
			return
		}
		content := fmt.Sprintf("Total join requests: %d\n", len(result.Data))
		for _, req := range result.Data {
			content += fmt.Sprintf("- Username: %s, User ID: %d, Request ID: %d\n", req.User.Name, req.User.Id, req.Id)
		}
		util.RespondEphemeral(s, i.Interaction, content)
	}
}
