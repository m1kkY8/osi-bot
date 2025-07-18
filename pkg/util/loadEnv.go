package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string, error) {
	_ = godotenv.Load()
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		return "", "", fmt.Errorf("HTB_TOKEN not set")
	}

	htbToken := os.Getenv("HTB_TOKEN")
	if discordToken == "" {
		return "", "", fmt.Errorf("HTB_TOKEN not set")
	}

	return discordToken, htbToken, nil
}
