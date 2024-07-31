package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WhatsAppWriter struct {
	apiURL string
	apiKey string
}

func NewWhatsAppWriter() *WhatsAppWriter {
	return &WhatsAppWriter{
		apiURL: "https://api.whatsapp.com/v1/messages",
		apiKey: "YOUR_WHATSAPP_API_KEY",
	}
}

func (w *WhatsAppWriter) WriteDraft(content string) error {
	message := map[string]interface{}{
		"to":   "whatsapp_number", // substitua pelo número do destinatário
		"type": "text",
		"text": map[string]string{"body": content},
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", w.apiURL, bytes.NewBuffer(messageBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+w.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API error: %v", resp.Status)
	}

	return nil
}
