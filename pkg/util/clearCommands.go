package util

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// ClearSlashCommands deletes all application (slash) commands from a guild or globally.
// Pass guildID for guild commands, or "" for global commands.
// Returns number of deleted commands and error if any.
func ClearSlashCommands(session *discordgo.Session, guildID string) (int, error) {
	appID := session.State.User.ID
	commands, err := session.ApplicationCommands(appID, guildID)
	if err != nil {
		return 0, fmt.Errorf("failed to list commands: %w", err)
	}
	count := 0
	for _, cmd := range commands {
		err := session.ApplicationCommandDelete(appID, guildID, cmd.ID)
		if err != nil {
			fmt.Printf("Failed to delete command '%s': %v\n", cmd.Name, err)
		} else {
			fmt.Printf("Deleted command: %s\n", cmd.Name)
			count++
		}
	}
	return count, nil
}
