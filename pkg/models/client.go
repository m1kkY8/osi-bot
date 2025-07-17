package models

import (
	"context"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gubarz/gohtb"
	"github.com/m1kkY8/osi-bot/pkg/bot/intents"
)

type Client struct {
	Context             context.Context
	HTBClient           gohtb.Client
	TeamMembers         []TeamMember
	BookstackUsers      []BookstackUser
	DiscordSession      *discordgo.Session
	ApplicationCommands []*discordgo.ApplicationCommand
	intents             []discordgo.Intent // Intents for the Discord session
	guildID             string             // empty for global, set for guild-specific commands
	adminRoleID         string             // Role ID for admin users, if needed
	teamID              string
}

func NewClient(teamMembers []TeamMember, discordSession *discordgo.Session) *Client {
	return &Client{
		TeamMembers:         teamMembers,
		BookstackUsers:      []BookstackUser{},
		DiscordSession:      discordSession,
		ApplicationCommands: []*discordgo.ApplicationCommand{},
		intents:             []discordgo.Intent{}, // Default to no intents
		guildID:             "",                   // Default to global commands
	}
}

func (c *Client) Initialize() {
	c.DiscordSession.Identify.Intents = intents.SetIntents()
	c.ApplicationCommands = setApplicationCommands()
	c.SetGuildID(os.Getenv("GUILD_ID"))
	c.SetAdminRoleID(os.Getenv("ADMIN_ROLE_ID"))
	c.SetTeamID(os.Getenv("HTB_TEAM_ID"))
}

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

func (c *Client) SetGuildID(guildID string) {
	c.guildID = guildID
}

func (c *Client) GetGuildID() string {
	return c.guildID
}

func (c *Client) SetTeamID(teamID string) {
	c.teamID = teamID
}

func (c *Client) GetTeamID() string {
	return c.teamID
}

func (c *Client) SetAdminRoleID(roleID string) {
	c.adminRoleID = roleID
}

func (c *Client) GetAdminRoleID() string {
	return c.adminRoleID
}

func (c *Client) SetIntents(intents []discordgo.Intent) {
	c.intents = intents
}

func (c *Client) GetIntents() []discordgo.Intent {
	return c.intents
}

func (c *Client) SetApplicationCommands(commands []*discordgo.ApplicationCommand) {
	c.ApplicationCommands = commands
}

func (c *Client) GetApplicationCommands() []*discordgo.ApplicationCommand {
	return c.ApplicationCommands
}
