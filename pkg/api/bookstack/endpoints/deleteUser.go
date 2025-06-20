package endpoints

import (
	"fmt"
	"io"
	"net/http"

	"github.com/m1kkY8/osi-bot/pkg/api/bookstack/auth"
	"github.com/m1kkY8/osi-bot/pkg/models"
)

func BookApiDeleteUser(id int) {
	url := fmt.Sprintf("%s/api/users/%d", models.BOOKSTACK_DOMAIN, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("[ERROR] Failed to create request:", err)
		return
	}

	for key, value := range auth.GetAuthHeader() {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[ERROR] Failed to fetch users:", err)
		return
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ERROR] Failed to read response body:", err)
		return
	}

	fmt.Println(rawBody)
}
