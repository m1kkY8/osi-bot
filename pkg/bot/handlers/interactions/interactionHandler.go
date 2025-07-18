package interactions

import (
	"github.com/m1kkY8/osi-bot/pkg/types"
)

func RegisterInteractionHandlers(client *types.Client, lbPages, bookstackPages *types.Page) {
	// Register user list interaction handler
	UserListInteraction(client, bookstackPages)

	// Register leaderboard interaction handler
	LeaderboardInteraction(client, lbPages)

	// Add more interaction handlers as needed
}
