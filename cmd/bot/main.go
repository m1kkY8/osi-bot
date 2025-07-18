package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gubarz/gohtb"
	"github.com/joho/godotenv"
	"github.com/m1kkY8/osi-bot/pkg/bot/handlers/commands"
	"github.com/m1kkY8/osi-bot/pkg/bot/handlers/interactions"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		fmt.Println("DISCORD_TOKEN not set")
		return
	}

	htbToken := os.Getenv("HTB_TOKEN")
	if discordToken == "" {
		fmt.Println("HTB_TOKEN not set")
		return
	}

	ctx := context.Background()
	htbClient, _ := gohtb.New(htbToken)

	// Create Discord session
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Printf("Error creating Discord session: %v\n", err)
		return
	}

	// Initialize stateful client and pagers
	client := models.NewClient(nil, dg)
	lbPages := models.NewPage(1, 10, 0, make(map[string]int))
	bookstackPages := models.NewPage(1, 10, 0, make(map[string]int))

	// Initialize client state (intents, commands, config)
	client.HTBClient = *htbClient
	client.Context = ctx

	models.InitializeClient(client)

	// Register custom interaction handlers (components, buttons, etc.)
	interactions.RegisterInteractionHandlers(client, lbPages, bookstackPages)

	// Register universal interaction dispatcher
	registerUniversalDispatcher(client, lbPages, bookstackPages)

	// Open Discord session
	if err := client.DiscordSession.Open(); err != nil {
		fmt.Printf("Error opening Discord session: %v\n", err)
		return
	}

	// Register slash commands
	client.RegisterSlashCommands()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait for graceful shutdown
	waitForInterrupt()
}

// registerUniversalDispatcher wires up the main dispatcher for slash and component interactions.
func registerUniversalDispatcher(client *models.Client, lbPages, bookstackPages *models.Page) {
	client.DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			commands.HandleSlashCommand(client, lbPages, bookstackPages, s, i)
		case discordgo.InteractionMessageComponent:
			fmt.Printf("[LOG] Component interaction: %s by %s\n", i.MessageComponentData().CustomID, i.Member.User.Username)
			// Actual handling occurs in RegisterInteractionHandlers
		}
	})
}

// waitForInterrupt blocks until CTRL+C or termination signal is received.
func waitForInterrupt() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("Shutting down gracefully...")
}
