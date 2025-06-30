package interactions

import "github.com/m1kkY8/osi-bot/pkg/models"

func RegisterInteractionHandlers(client *models.Client, lbPages, bookstackPages *models.Page) {
	// Register user list interaction handler
	UserListInteraction(client, bookstackPages)

	// Register leaderboard interaction handler
	LeaderboardInteraction(client, lbPages)

	// Add more interaction handlers as needed
}
