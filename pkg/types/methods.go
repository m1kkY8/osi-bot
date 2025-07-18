package types

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Client methods

// Initialize sets up the client with environment variables and default values
func (c *Client) Initialize() {
	// Note: intents and setApplicationCommands need to be called from models package
	// to avoid circular imports
	c.SetGuildID(os.Getenv("GUILD_ID"))
	c.SetAdminRoleID(os.Getenv("ADMIN_ROLE_ID"))
	c.SetTeamID(os.Getenv("HTB_TEAM_ID"))
}

// RegisterSlashCommands registers all application commands with Discord
func (c *Client) RegisterSlashCommands() {
	for _, cmd := range c.ApplicationCommands {
		_, err := c.DiscordSession.ApplicationCommandCreate(
			c.DiscordSession.State.User.ID,
			c.GetGuildID(),
			cmd,
		)
		if err != nil {
			fmt.Printf("Error creating command '%s': %v\n", cmd.Name, err)
		} else {
			fmt.Printf("Registered command: %s\n", cmd.Name)
		}
	}
}

// Getter and Setter methods
func (c *Client) SetGuildID(guildID string) {
	c.GuildID = guildID
}

func (c *Client) GetGuildID() string {
	return c.GuildID
}

func (c *Client) SetTeamID(teamID string) {
	c.TeamID = teamID
}

func (c *Client) GetTeamID() string {
	return c.TeamID
}

func (c *Client) SetAdminRoleID(roleID string) {
	c.AdminRoleID = roleID
}

func (c *Client) GetAdminRoleID() string {
	return c.AdminRoleID
}

func (c *Client) SetIntents(intents []discordgo.Intent) {
	c.Intents = intents
}

func (c *Client) GetIntents() []discordgo.Intent {
	return c.Intents
}

func (c *Client) SetApplicationCommands(commands []*discordgo.ApplicationCommand) {
	c.ApplicationCommands = commands
}

func (c *Client) GetApplicationCommands() []*discordgo.ApplicationCommand {
	return c.ApplicationCommands
}
