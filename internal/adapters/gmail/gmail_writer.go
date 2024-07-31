package gmail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marquesmarcelo/marcelo_iapkg/httpclient"
)

type GmailWriter struct {
	client *httpclient.HTTPClient
}

func NewGmailWriter(apiKey string) *GmailWriter {
	return &GmailWriter{
		client: httpclient.New(apiKey),
	}
}

func (w *GmailWriter) WriteDraft(content string) error {
	url := "https://gmail.googleapis.com/gmail/v1/users/me/drafts"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	message := map[string]interface{}{
		"message": map[string]string{
			"raw": content,
		},
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
