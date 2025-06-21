package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
)

func DeleteUserSlashHandler(adminRoleID string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		hasAdmin := slices.Contains(i.Member.Roles, adminRoleID)
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
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "❌ You must provide a valid user ID.",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}

		endpoints.BookApiDeleteUser(userID) // adjust type if needed
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("✅ Attempting to delete user with ID: %s", userID),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}
