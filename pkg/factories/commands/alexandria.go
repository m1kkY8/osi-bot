package commandFactory

import "github.com/bwmarrin/discordgo"

// CreateRegisterCommand creates the register command for Alexandria
func CreateRegisterCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "register",
		Description: "Register a user in Alexandria",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "username",
				Description: "The username to register",
				Required:    true,
			},
		},
	}
}

// CreateUpdateCommand creates the update command for Alexandria
func CreateUpdateCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "update",
		Description: "Change users permission level",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "user_id",
				Description: "The user ID to promote",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "role",
				Description: "Permission level (viewer/editor)",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "viewer",
						Value: "viewer",
					},
					{
						Name:  "editor",
						Value: "editor",
					},
				},
			},
		},
	}
}

// CreateRemoveCommand creates the remove command for Alexandria
func CreateRemoveCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "remove",
		Description: "Remove a user from Alexandria",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "user_id",
				Description: "The user ID to remove",
				Required:    true,
			},
		},
	}
}

// CreateUsersCommand creates the users command for Alexandria
func CreateUsersCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "users",
		Description: "List all users in Alexandria",
	}
}

// GetAllAlexandriaCommands returns all Alexandria commands
func GetAllAlexandriaCommands() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		CreateRegisterCommand(),
		CreateUpdateCommand(),
		CreateRemoveCommand(),
		CreateUsersCommand(),
	}
}
