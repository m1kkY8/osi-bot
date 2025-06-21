package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/auth"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/endpoints"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

// Pass adminRoleID as string argument when registering the handler
func RegisterUserSlashHandler(client *models.Client, adminRoleID string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Check for admin role
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
		var targetUser *discordgo.User
		if len(options) > 0 && options[0].Name == "register" && options[0].Type == discordgo.ApplicationCommandOptionSubCommand {
			for _, opt := range options[0].Options {
				if opt.Name == "username" && opt.Type == discordgo.ApplicationCommandOptionUser {
					targetUser = opt.UserValue(s)
					break
				}
			}
		}

		if targetUser == nil {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "❌ Could not find the specified user.",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}

		// Generate credentials for the target user
		username := strings.ToLower(targetUser.Username)
		email := fmt.Sprintf("%s@offsecinitiative.net", username)
		password := auth.GeneratePassword()

		bookUser := models.CreateBookstackUser(username, email, password)
		statusCode, _ := endpoints.BookApiCreateUser(bookUser)

		var dmMessage string
		if statusCode == 422 {
			dmMessage = "❌ User already exists. Please try a different username."
		} else if statusCode >= 300 {
			dmMessage = "❌ Failed to create BookStack user."
		} else {
			dmMessage = fmt.Sprintf(
				"✅ BookStack user created for you!\n**Username:** %s\n**Email:** %s\n**Password:** %s",
				username, email, password,
			)
		}

		// DM the user if possible
		dmSent := false
		if statusCode < 300 {
			if dmChannel, err := s.UserChannelCreate(targetUser.ID); err == nil {
				if _, err := s.ChannelMessageSend(dmChannel.ID, dmMessage); err == nil {
					dmSent = true
				}
			}
		}

		// Acknowledge the admin
		var response string
		if statusCode < 300 && dmSent {
			response = fmt.Sprintf("Registration successful! Credentials sent via DM to <@%s>.", targetUser.ID)
		} else if statusCode < 300 {
			response = fmt.Sprintf("Registration successful, but failed to DM <@%s> (maybe DMs are blocked).", targetUser.ID)
		} else {
			response = fmt.Sprintf("Failed to register user: %s", dmMessage)
		}
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}
