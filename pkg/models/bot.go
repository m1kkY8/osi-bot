package models

import "github.com/bwmarrin/discordgo"

type Client struct {
	TeamMembers         []TeamMember
	BookstackUsers      []BookstackUser
	DiscordSession      *discordgo.Session
	ApplicationCommands []*discordgo.ApplicationCommand
	intents             []discordgo.Intent // Intents for the Discord session
	guildID             string             // empty for global, set for guild-specific commands
	adminRoleID         string             // Role ID for admin users, if needed
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

func (c *Client) SetGuildID(guildID string) {
	c.guildID = guildID
}

func (c *Client) GetGuildID() string {
	return c.guildID
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
