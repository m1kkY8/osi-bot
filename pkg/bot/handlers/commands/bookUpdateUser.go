package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
	"github.com/m1kkY8/osi-bot/pkg/types"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func updateUserSlashHandler(client *types.Client) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Check for admin role
		hasAdmin := slices.Contains(i.Member.Roles, client.GetAdminRoleID())
		if !hasAdmin {
			util.RespondEphemeral(s, i.Interaction, "❌ You do not have permission to use this command.")
			return
		}

		options := i.ApplicationCommandData().Options
		var userID string
		var newRole string

		if len(options) > 0 && options[0].Name == "update" && options[0].Type == discordgo.ApplicationCommandOptionSubCommand {
			for _, opt := range options[0].Options {
				if opt.Name == "user_id" && opt.Type == discordgo.ApplicationCommandOptionString {
					userID = opt.StringValue()
				}
				if opt.Name == "role" && opt.Type == discordgo.ApplicationCommandOptionString {
					newRole = opt.StringValue()
				}
			}
		}

		if userID == "" {
			util.RespondEphemeral(s, i.Interaction, "❌ No BookStack user ID provided.")
			return
		}
		if newRole == "" {
			util.RespondEphemeral(s, i.Interaction, "❌ No role specified for update (viewer/editor).")
			return
		}

		// Compose minimal update object
		updatedUser := &types.BookstackUser{}
		var roleID int
		switch strings.ToLower(newRole) {
		case "editor":
			roleID = types.EDITOR
		case "viewer":
			roleID = types.VIEWER
		default:
			util.RespondEphemeral(s, i.Interaction, "❌ Invalid role. Choose either 'editor' or 'viewer'.")
			return
		}
		updatedUser.Roles = []int{roleID}

		statusCode, _ := endpoints.BookApiUpdateUser(userID, updatedUser)

		var response string
		if statusCode < 300 {
			response = fmt.Sprintf("✅ BookStack user with ID `%s` updated to role: %s.", userID, newRole)
		} else if statusCode == 404 {
			response = "❌ BookStack user not found for update."
		} else {
			response = "❌ Failed to update BookStack user."
		}
		util.RespondEphemeral(s, i.Interaction, response)
	}
}
