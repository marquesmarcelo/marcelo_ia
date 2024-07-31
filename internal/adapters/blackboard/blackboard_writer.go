package blackboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marquesmarcelo/marcelo_iapkg/httpclient"
)

type BlackboardWriter struct {
	client *httpclient.HTTPClient
}

func NewBlackboardWriter(apiKey string) *BlackboardWriter {
	return &BlackboardWriter{
		client: httpclient.New(apiKey),
	}
}

func (w *BlackboardWriter) WriteDraft(content string) error {
	url := "https://blackboard.example.com/learn/api/public/v1/messages"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	message := map[string]interface{}{
		"body": content,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := w.client.Post(url, bytes.NewBuffer(jsonData), headers)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API error: %v", resp.Status)
	}

	return nil
}
