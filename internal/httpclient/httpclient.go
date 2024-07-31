package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
	APIKey string
}

func New(apiKey string) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		APIKey: apiKey,
	}
}

func (c *HTTPClient) Get(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *HTTPClient) Post(url string, body interface{}, headers map[string]string) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

func (c *HTTPClient) Patch(url string, body interface{}, headers map[string]string) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}
