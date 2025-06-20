package auth

import (
	"fmt"
	"os"
)

func GetAuthHeader() map[string]string {
	tokenId := os.Getenv("BOOKSTACK_ID")
	tokenSecret := os.Getenv("BOOKSTACK_TOKEN")
	auth := fmt.Sprintf("Token %s:%s", tokenId, tokenSecret)

	return map[string]string{
		"Authorization": auth,
		"Content-Type":  "application/json",
	}
}
