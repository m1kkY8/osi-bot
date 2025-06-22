package htb

import "os"

var (
	BaseURL = "https://labs.hackthebox.com/api/v4"
	TeamID  = "7185"
)

// Structs for API responses
type JoinRequest struct {
	ID   int `json:"id"`
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}

type JoinRequestsResponse struct {
	Original []JoinRequest `json:"original"`
}

func getHTBToken() string {
	return os.Getenv("HTB_TOKEN")
}
