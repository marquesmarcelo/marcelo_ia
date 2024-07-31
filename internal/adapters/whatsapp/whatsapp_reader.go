package whatsapp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marquesmarcelo/marcelo_ia/domain"
)

type WhatsAppReader struct {
	apiURL string
	apiKey string
}

func NewWhatsAppReader() *WhatsAppReader {
	return &WhatsAppReader{
		apiURL: "https://api.whatsapp.com/v1/messages",
		apiKey: "YOUR_WHATSAPP_API_KEY",
	}
}

func (r *WhatsAppReader) ReadUnreadMessages() ([]domain.Message, error) {
	req, err := http.NewRequest("GET", r.apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+r.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %v", resp.Status)
	}

	var result struct {
		Messages []domain.Message `json:"messages"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Messages, nil
}
