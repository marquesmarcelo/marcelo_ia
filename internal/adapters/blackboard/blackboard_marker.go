package blackboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marquesmarcelo/marcelo_iapkg/httpclient"
)

type BlackboardMarker struct {
	client *httpclient.HTTPClient
}

func NewBlackboardMarker(apiKey string) *BlackboardMarker {
	return &BlackboardMarker{
		client: httpclient.New(apiKey),
	}
}

func (m *BlackboardMarker) MarkMessageAsRead(messageID string) error {
	url := fmt.Sprintf("https://blackboard.example.com/learn/api/public/v1/messages/%s", messageID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	payload := map[string]interface{}{
		"status": "read",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := m.client.Patch(url, bytes.NewBuffer(jsonData), headers)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %v", resp.Status)
	}

	return nil
}
