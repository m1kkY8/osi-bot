package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack"
	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/auth"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func RegisterUser(client *models.Client) {
	client.DiscordSession.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot || !strings.HasPrefix(m.Content, "!register") {
			return
		}

		fmt.Printf("[INFO] Received !register command from user: %s\n", m.Author.Username)

		username := strings.ToLower(m.Author.Username)
		email := fmt.Sprintf("%s@offsecinitiative.net", username)
		password := auth.GeneratePassword()

		fmt.Printf("[DEBUG] Generated credentials - Username: %s, Email: %s, Password: %s\n", username, email, password)

		user := bookstack.CreateBookstackUser(username, email, password)

		body, err := json.Marshal(user)
		if err != nil {
			fmt.Println("[ERROR] Failed to marshal user JSON:", err)
			return
		}
		fmt.Println("[DEBUG] Marshaled user payload:", string(body))

		req, err := http.NewRequest("POST", models.BOOKSTACK_DOMAIN+"/api/users", strings.NewReader(string(body)))
		if err != nil {
			fmt.Println("[ERROR] Failed to create HTTP request:", err)
			return
		}

		authHeaders := auth.GetAuthHeader()
		for key, value := range authHeaders {
			req.Header.Add(key, value)
			fmt.Printf("[DEBUG] Set HTTP header: %s: %s\n", key, value)
		}

		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("[ERROR] HTTP request failed:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Printf("[INFO] BookStack API responded with status: %d\n", resp.StatusCode)
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("[DEBUG] API response body: %s\n", string(respBody))

		if resp.StatusCode == 422 {
			fmt.Println("[ERROR] User already exists. Please try a different username.")
			dm, err := client.DiscordSession.UserChannelCreate(m.Author.ID)
			if err != nil {
				fmt.Println("[ERROR] Error creating DM channel:", err)
				return
			}

			_, err = client.DiscordSession.ChannelMessageSend(dm.ID, "❌ User already exists. Please try a different username.")
			if err != nil {
				fmt.Println("[ERROR] Error sending DM message:", err)
				return
			}
			return
		}

		if resp.StatusCode >= 300 {
			fmt.Println("[ERROR] Failed to create BookStack user.")
			return
		}

		dm, err := client.DiscordSession.UserChannelCreate(m.Author.ID)
		if err != nil {
			fmt.Println("[ERROR] Error creating DM channel:", err)
			return
		}

		message := fmt.Sprintf("✅ User created successfully!\n**Username:** %s\n**Email:** %s\n**Password:** %s", username, email, password)
		_, err = client.DiscordSession.ChannelMessageSend(dm.ID, message)
		if err != nil {
			fmt.Println("[ERROR] Error sending DM message:", err)
			return
		}

		fmt.Printf("[INFO] Sent credentials to %s via DM.\n", m.Author.Username)
	})
}
