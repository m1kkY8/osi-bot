package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/m1kkY8/osi-bot/pkg/bot/handlers"
	"github.com/m1kkY8/osi-bot/pkg/intents"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func main() {
	godotenv.Load()

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	client := models.NewClient(nil, dg)
	lbPages := models.NewPage(1, 10, 0, make(map[string]int))
	bookstackPages := models.NewPage(1, 10, 0, make(map[string]int))

	client.DiscordSession.Identify.Intents = intents.SetIntents()

	handlers.LeaderboardHandler(client, lbPages)
	handlers.LeaderboardInteraction(client, lbPages)
	handlers.RegisterUser(client)
	handlers.GetUsers(client, bookstackPages)
	handlers.UserListInteraction(client, bookstackPages)
	handlers.DeleteUser(client)

	err = client.DiscordSession.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
