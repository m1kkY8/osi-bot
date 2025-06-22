package htb

import (
	"fmt"
	"io"
	"net/http"
)

// Kick a user from the team
func HTBKickUser(userID string) error {
	url := fmt.Sprintf("%s/team/kick/%s", BaseURL, userID)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", "Bearer "+getHTBToken())
	req.Header.Set("User-Agent", "curl/7.68.0")
	req.Header.Set("accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s", string(b))
	}
	return nil
}
