package auth

import (
	"fmt"

	"github.com/m1kkY8/osi-bot/pkg/util"
)

func GetAuthHeader() map[string]string {
	tokenSecret, tokenId, err := util.LoadAuthEnv()

	if err != nil {
		fmt.Printf("Error loading BookStack environment variables: %v\n", err)
		return nil
	}

	auth := fmt.Sprintf("Token %s:%s", tokenId, tokenSecret)

	return map[string]string{
		"Authorization": auth,
		"Content-Type":  "application/json",
	}
}
