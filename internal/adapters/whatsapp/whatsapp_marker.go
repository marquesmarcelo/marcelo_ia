package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WhatsAppMarker struct {
	apiURL string
	apiKey string
}

func NewWhatsAppMarker() *WhatsAppMarker {
	return &WhatsAppMarker{
		apiURL: "https://api.whatsapp.com/v1/messages",
		apiKey: "YOUR_WHATSAPP_API_KEY",
	}
}

func (m *WhatsAppMarker) MarkMessageAsRead(messageID string) error {
	url := fmt.Sprintf("%s/%s", m.apiURL, messageID)
	payload := map[string]interface{}{
		"status": "read",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %v", resp.Status)
	}

	return nil
}
