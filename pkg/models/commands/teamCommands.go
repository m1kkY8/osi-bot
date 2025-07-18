package models

import "github.com/bwmarrin/discordgo"

func GetTeamCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "team",
			Description: "Manage HTB Team",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "invitations",
					Description: "Get all join requests",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "accept",
					Description: "Accept a join request",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "request_id",
							Description: "Request ID to accept",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "reject",
					Description: "Reject a join request",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "request_id",
							Description: "Request ID to reject",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "kick",
					Description: "Kick a user from the team",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "user_id",
							Description: "User ID to kick",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "leaderboard",
					Description: "Show the leaderboard",
				},
			},
		},
	}
}
