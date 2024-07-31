package gmail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marquesmarcelo/marcelo_iapkg/httpclient"
)

type GmailMarker struct {
	client *httpclient.HTTPClient
}

func NewGmailMarker(apiKey string) *GmailMarker {
	return &GmailMarker{
		client: httpclient.New(apiKey),
	}
}

func (m *GmailMarker) MarkMessageAsRead(messageID string) error {
	url := fmt.Sprintf("https://gmail.googleapis.com/gmail/v1/users/me/messages/%s/modify", messageID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	payload := map[string]interface{}{
		"removeLabelIds": []string{"UNREAD"},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := m.client.Post(url, bytes.NewBuffer(jsonData), headers)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %v", resp.Status)
	}

	return nil
}
