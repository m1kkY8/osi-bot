package htb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"

	"github.com/m1kkY8/osi-bot/pkg/models"
)

func FetchUsers() ([]models.TeamMember, error) {
	apiURL := "https://www.hackthebox.com/api/v4/team/members/6859"
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("User-Agent", "curl/7.64.1")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("HTB_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch users, status code:", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch users, status code: %d", resp.StatusCode)
	}

	var parsed []models.TeamMember
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	sort.Slice(parsed, func(i, j int) bool {
		return parsed[i].Points > parsed[j].Points
	})

	return parsed, nil
}
