package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/auth"
	"github.com/m1kkY8/osi-bot/pkg/models"
	"github.com/m1kkY8/osi-bot/pkg/models/roles"
)

// Update an existing BookStack user by ID
func BookApiUpdateUser(id string, user *models.BookstackUser) (int, error) {
	body, err := json.Marshal(user)
	if err != nil {
		fmt.Println("[ERROR] Failed to marshal user JSON:", err)
		return -1, err
	}

	url := fmt.Sprintf("%s/api/users/%s", models.BOOKSTACK_DOMAIN, id)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("[ERROR] Failed to create HTTP request:", err)
		return -1, err
	}

	authHeaders := auth.GetAuthHeader()
	for key, value := range authHeaders {
		req.Header.Add(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("[ERROR] HTTP request failed:", err)
		return -1, err
	}
	defer resp.Body.Close()

	fmt.Printf("[INFO] BookStack API responded with status: %d\n", resp.StatusCode)
	return resp.StatusCode, nil
}

// Promote a BookStack user from viewer to editor role
func PromoteBookstackUserToEditor(existingUser *models.BookstackUser) *models.BookstackUser {
	user := *existingUser // Copy to avoid mutating original
	user.Roles = []int{roles.EDITOR}
	return &user
}
