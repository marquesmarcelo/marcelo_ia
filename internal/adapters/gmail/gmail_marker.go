package gmail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marquesmarcelo/marcelo_ia/pkg/httpclient"
)

type GmailMarker struct {
	client *httpclient.HTTPClient
	apiUrl string
}

func NewGmailMarker(client *httpclient.HTTPClient, apiUrl string) *GmailMarker {
	return &GmailMarker{
		client: client,
		apiUrl: apiUrl,
	}
}

func (m *GmailMarker) MarkMessageAsRead(messageID string) error {
	url := fmt.Sprintf("%s/users/me/messages/%s/modify", m.apiUrl, messageID)
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
