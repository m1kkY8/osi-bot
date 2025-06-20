package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/auth"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func BookApiCreateUser(user *models.BookstackUser) (int, error) {
	body, err := json.Marshal(user)
	if err != nil {
		fmt.Println("[ERROR] Failed to marshal user JSON:", err)
		return -1, err
	}
	fmt.Println("[DEBUG] Marshaled user payload:", string(body))

	req, err := http.NewRequest("POST", models.BOOKSTACK_DOMAIN+"/api/users", strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("[ERROR] Failed to create HTTP request:", err)
		return -1, err
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
		return -1, err
	}
	defer resp.Body.Close()

	fmt.Printf("[INFO] BookStack API responded with status: %d\n", resp.StatusCode)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("[DEBUG] API response body: %s\n", string(respBody))

	return resp.StatusCode, nil
}
