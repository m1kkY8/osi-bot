package endpoints

import (
	"fmt"
	"io"
	"net/http"

	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/auth"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func BookApiDeleteUser(id string) error {
	url := fmt.Sprintf("%s/api/users/%s", models.BOOKSTACK_DOMAIN, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("[ERROR] Failed to create request:", err)
		return err
	}

	for key, value := range auth.GetAuthHeader() {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[ERROR] Failed to fetch users:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("[INFO] User not found or already deleted")
		return nil // treat as success
	}

	if resp.StatusCode == 500 {
		fmt.Println("[ERROR] Internal server error")
		return fmt.Errorf("internal server error while deleting user with ID %s", id)
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ERROR] Failed to read response body:", err)
		return err
	}

	fmt.Println(string(rawBody))
	return nil
}
