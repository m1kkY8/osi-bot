package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/htb"
	"github.com/m1kkY8/osi-bot/pkg/models"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func teamRejectSlashHandler(client *models.Client) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {
			util.RespondEphemeral(s, i.Interaction, "❌ You do not have permission to use this command.")
			return
		}
		sub := i.ApplicationCommandData().Options[0]
		var requestID string
		for _, opt := range sub.Options {
			if opt.Name == "request_id" {
				requestID = opt.StringValue()
			}
		}
		if requestID == "" {
			util.RespondEphemeral(s, i.Interaction, "Missing request ID.")
			return
		}
		err := htb.HTBAcceptJoin(requestID)
		if err == nil {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("✅ Successfully rejected invite for request ID %s", requestID))
		} else {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("❌ Failed to reject invite for request ID %s: %v", requestID, err))
		}
	}
}
