package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
	"github.com/m1kkY8/osi-bot/pkg/types"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func deleteUserSlashHandler(client *types.Client) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {
			util.RespondEphemeral(s, i.Interaction, "❌ You do not have permission to use this command.")
			return
		}

		options := i.ApplicationCommandData().Options
		var userID string
		if len(options) > 0 && options[0].Name == "remove" && options[0].Type == discordgo.ApplicationCommandOptionSubCommand {
			for _, opt := range options[0].Options {
				if opt.Name == "user_id" && opt.Type == discordgo.ApplicationCommandOptionString {
					userID = opt.StringValue()
				}
			}
		}
		if userID == "" {
			util.RespondEphemeral(s, i.Interaction, "❌ You must provide a valid user ID.")
			return
		}

		endpoints.BookApiDeleteUser(userID) // adjust type if needed
		if err := endpoints.BookApiDeleteUser(userID); err != nil {
			util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("❌ Failed to delete user with ID: %s. Error: %s", userID, err.Error()))
			return
		}
		util.RespondEphemeral(s, i.Interaction, fmt.Sprintf("✅ Deleted user with ID: %s", userID))
	}
}
