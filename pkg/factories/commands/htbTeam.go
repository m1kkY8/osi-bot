package commandFactory

import "github.com/bwmarrin/discordgo"

// CreateInvitationsCommand creates the invitations command for HTB Team
func CreateInvitationsCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "invitations",
		Description: "Get all join requests",
	}
}

// CreateAcceptCommand creates the accept command for HTB Team
func CreateAcceptCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
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
	}
}

// CreateRejectCommand creates the reject command for HTB Team
func CreateRejectCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
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
	}
}

// CreateKickCommand creates the remove user command for HTB Team
func CreateKickCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
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
	}
}

// CreateLeaderboardCommand creates the users command for HTB Team
func CreateLeaderboardCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "leaderboard",
		Description: "Show the leaderboard",
	}
}

// GetAllHTBTeamCommands returns all HTB Team commands
func GetAllHTBTeamCommands() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		CreateInvitationsCommand(),
		CreateRejectCommand(),
		CreateKickCommand(),
		CreateLeaderboardCommand(),
		CreateAcceptCommand(),
	}
}
