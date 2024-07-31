package gmail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marquesmarcelo/marcelo_ia/internal/httpclient"
)

type GmailWriter struct {
	client *httpclient.HTTPClient
	apiUrl string
}

func NewGmailWriter(client *httpclient.HTTPClient, apiUrl string) *GmailWriter {
	return &GmailWriter{
		client: client,
		apiUrl: apiUrl,
	}
}

func (w *GmailWriter) WriteDraft(content string) error {
	url := fmt.Sprintf("%s/users/me/drafts", w.apiUrl)
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
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API error: %v", resp.Status)
	}

	return nil
}
