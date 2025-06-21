package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
)

func DeleteUserSlashHandler(adminRoleID string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand || i.ApplicationCommandData().Name != "remove" {
			return
		}

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

		// Get the user_id argument
		options := i.ApplicationCommandData().Options
		var userID int
		for _, opt := range options {
			if opt.Name == "user_id" && opt.Type == discordgo.ApplicationCommandOptionInteger {
				userID = int(opt.IntValue())
			}
		}
		if userID == 0 {
			// Should not happen if Required: true, but safety check
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "❌ You must provide a valid user ID.",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}

		endpoints.BookApiDeleteUser(userID)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("✅ Attempting to delete user with ID: %d", userID),
				Flags:   discordgo.MessageFlagsEphemeral, // Use Flags: discordgo.MessageFlagsEphemeral for private response
			},
		})
	}
}
