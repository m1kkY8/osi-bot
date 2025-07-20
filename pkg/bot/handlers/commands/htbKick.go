package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/types"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func teamKickSlashHandler(client *types.Client) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {
			_ = util.RespondEphemeral(s, i.Interaction, "❌ You do not have permission to use this command.")
			return
		}
		sub := i.ApplicationCommandData().Options[0]
		var userID int
		for _, opt := range sub.Options {
			if opt.Name == "user_id" {
				userID = int(opt.IntValue())
			}
		}
		if userID == 0 {
			util.RespondEphemeral(s, i.Interaction, "Missing user ID.")
			return
		}
		restult, err := client.HTBClient.Teams.KickMember(client.Context, userID)
		fmt.Println(restult.Data.Message)
		if err == nil {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("✅ Kicked user ID %d from the team.", userID))
		} else {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("❌ Failed to kick user ID %d: %v", userID, err))
		}
	}
}
