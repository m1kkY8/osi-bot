package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/m1kkY8/osi-bot/pkg/bot/handlers/commands"
	"github.com/m1kkY8/osi-bot/pkg/bot/handlers/interactions"
	"github.com/m1kkY8/osi-bot/pkg/bot/intents"
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

	// Set up intents and application commands
	client.DiscordSession.Identify.Intents = intents.SetIntents()
	client.ApplicationCommands = models.SetApplicationCommands()
	client.SetGuildID(os.Getenv("GUILD_ID"))
	client.SetAdminRoleID(os.Getenv("ADMIN_ROLE_ID"))

	// Register interaction handlers (buttons etc)
	interactions.UserListInteraction(client, bookstackPages)
	interactions.LeaderboardInteraction(client, lbPages)

	// Slash command handler map using factories; pass client/pages as needed
	slashHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"leaderboard": commands.LeaderboardSlashHandler(client, lbPages),
	}

	// Universal slash command dispatcher
	client.DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			data := i.ApplicationCommandData()
			if data.Name == "alexandria" && len(data.Options) > 0 {
				sub := data.Options[0]
				var handler func(*discordgo.Session, *discordgo.InteractionCreate)
				switch sub.Name {
				case "register":
					fmt.Printf("[LOG] /alexandria register called by %s\n", i.Member.User.Username)
					handler = commands.RegisterUserSlashHandler(client, client.GetAdminRoleID())
				case "remove":
					fmt.Printf("[LOG] /alexandria remove called by %s\n", i.Member.User.Username)
					handler = commands.DeleteUserSlashHandler(client.GetAdminRoleID())
				case "users":
					fmt.Printf("[LOG] /alexandria users called by %s\n", i.Member.User.Username)
					handler = commands.BookUserSlashCommandHandler(client, bookstackPages)
				default:
					fmt.Printf("[LOG] /alexandria unknown subcommand: %s by %s\n", sub.Name, i.Member.User.Username)
				}
				if handler != nil {
					handler(s, i)
				}
			} else if handler, ok := slashHandlers[data.Name]; ok {
				fmt.Printf("[LOG] /%s called by %s\n", data.Name, i.Member.User.Username)
				handler(s, i)
			} else {
				fmt.Printf("[LOG] Unknown command: %s by %s\n", data.Name, i.Member.User.Username)
			}
		case discordgo.InteractionMessageComponent:
			fmt.Printf("[LOG] Component interaction: %s by %s\n", i.MessageComponentData().CustomID, i.Member.User.Username)
		}
	})

	// Open the Discord session
	err = client.DiscordSession.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	// Slash commands registration
	for _, cmd := range client.ApplicationCommands {
		_, err := client.DiscordSession.ApplicationCommandCreate(
			client.DiscordSession.State.User.ID,
			client.GetGuildID(),
			cmd,
		)
		fmt.Printf("Registered command: %s\n", cmd.Name)
		if err != nil {
			fmt.Printf("Error creating command '%s': %v\n", cmd.Name, err)
		}
	}

	// deleted, err := util.ClearSlashCommands(client.DiscordSession, "1154887554965962932")
	// fmt.Printf("Deleted %d commands\n", deleted)

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
