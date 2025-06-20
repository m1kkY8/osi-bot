package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/m1kkY8/osi-bot/pkg/bot/handlers/commands"
	"github.com/m1kkY8/osi-bot/pkg/bot/handlers/interactions"
	"github.com/m1kkY8/osi-bot/pkg/intents"
	"github.com/m1kkY8/osi-bot/pkg/models"
	_ "github.com/m1kkY8/osi-bot/pkg/util"
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

	// Register interaction handlers (buttons etc)
	interactions.UserListInteraction(client, bookstackPages)
	interactions.LeaderboardInteraction(client, lbPages)

	// Slash command handler map using factories; pass client/pages as needed
	slashHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"alexandriaregister": commands.RegisterUserSlashHandler(client, "1356379340253958315"),
		"removeuser":         commands.DeleteUserSlashHandler("1356379340253958315"),
		"alexandriausers":    commands.BookUserSlashCommandHandler(client, bookstackPages),
		"leaderboard":        commands.LeaderboardSlashHandler(client, lbPages),
	}

	// Universal slash command dispatcher
	client.DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			if handler, ok := slashHandlers[i.ApplicationCommandData().Name]; ok {
				handler(s, i)
			}
		}
	})

	// Open the Discord session
	err = client.DiscordSession.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	// Slash commands registration
	applicationCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "alexandriaregister",
			Description: "Register a new BookStack user for a Discord user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The Discord user to register",
					Required:    true,
				},
			},
		},
		{
			Name:        "removeuser",
			Description: "Delete a BookStack user by their user ID",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "user_id",
					Description: "The BookStack user ID to delete",
					Required:    true,
				},
			},
		},
		{
			Name:        "alexandriausers",
			Description: "Get a list of BookStack users",
		},
		{
			Name:        "leaderboard",
			Description: "Show the leaderboard",
		},
	}

	for _, cmd := range applicationCommands {
		_, err := client.DiscordSession.ApplicationCommandCreate(
			client.DiscordSession.State.User.ID,
			"1154887554965962932",
			//			"1180277331189842021", // "" for global, or guild ID for quick updates
			cmd,
		)
		if err != nil {
			fmt.Printf("Error creating command '%s': %v\n", cmd.Name, err)
		}
	}

	// deleted, err := util.ClearSlashCommands(client.DiscordSession, "1154887554965962932")
	// fmt.Printf("Deleted %d commands\n", deleted)

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
