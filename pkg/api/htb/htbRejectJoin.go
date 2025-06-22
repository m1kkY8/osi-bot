package htb

import (
	"fmt"
	"io"
	"net/http"
)

func HTBRejectJoin(requestID string) error {
	// Construct the URL for accepting a join request

	url := fmt.Sprintf("%s/team/%s/invite/reject", BaseURL, requestID)
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Authorization", "Bearer "+getHTBToken())
	req.Header.Add("User-Agent", "curl/7.64.1")

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
