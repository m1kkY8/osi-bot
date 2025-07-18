package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/types"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func teamRejectSlashHandler(client *types.Client) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {
			util.RespondEphemeral(s, i.Interaction, "❌ You do not have permission to use this command.")
			return
		}
		sub := i.ApplicationCommandData().Options[0]
		var requestID int
		for _, opt := range sub.Options {
			if opt.Name == "request_id" {
				requestID = int(opt.IntValue())
			}
		}
		if requestID == 0 {
			util.RespondEphemeral(s, i.Interaction, "Missing request ID.")
			return
		}
		_, err := client.HTBClient.Teams.RejectInvite(client.Context, requestID)
		if err == nil {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("✅ Successfully rejected invite for request ID %d", requestID))
		} else {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("❌ Failed to reject invite for request ID %d: %v", requestID, err))
		}
	}
}
