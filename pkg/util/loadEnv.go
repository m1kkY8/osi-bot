package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadAuthEnv() (string, string, error) {
	_ = godotenv.Load()

	tokenId := os.Getenv("BOOKSTACK_ID")
	if tokenId == "" {
		return "", "", fmt.Errorf("BOOKSTACK_ID not set")
	}

	tokenSecret := os.Getenv("BOOKSTACK_TOKEN")
	if tokenSecret == "" {
		return "", "", fmt.Errorf("BOOKSTACK_TOKEN not set")
	}

	return tokenId, tokenSecret, nil
}

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
