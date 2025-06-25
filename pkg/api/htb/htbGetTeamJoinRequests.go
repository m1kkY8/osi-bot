package htb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetJoinRequests(teamID string) ([]JoinRequest, error) {
	url := fmt.Sprintf("%s/team/invitations/%s", BaseURL, teamID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+getHTBToken())
	req.Header.Set("User-Agent", "curl/7.68.0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(b))
	}
	var data JoinRequestsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	fmt.Printf("Retrieved %d join requests\n", len(data.Original))
	return data.Original, nil
}
