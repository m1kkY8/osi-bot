package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/auth"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func BookApiListUsers() []models.BookstackUser {
	req, err := http.NewRequest("GET", models.BOOKSTACK_DOMAIN+"/api/users", nil)
	if err != nil {
		fmt.Println("[ERROR] Failed to create request:", err)
		return nil
	}

	for key, value := range auth.GetAuthHeader() {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[ERROR] Failed to fetch users:", err)
		return nil
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ERROR] Failed to read response body:", err)
		return nil
	}

	var response models.BookstackUserResponse
	if err := json.Unmarshal(rawBody, &response); err != nil {
		fmt.Println("[ERROR] Failed to unmarshal users JSON:", err)
		return nil
	}

	for _, user := range response.Data {
		fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	return response.Data
}
