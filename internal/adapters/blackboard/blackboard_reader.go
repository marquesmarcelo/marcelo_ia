package blackboard

import (
	"encoding/json"

	"github.com/marquesmarcelo/marcelo_ia/domain"
	"github.com/marquesmarcelo/marcelo_iapkg/httpclient"
)

type BlackboardReader struct {
	client *httpclient.HTTPClient
}

func NewBlackboardReader(apiKey string) *BlackboardReader {
	return &BlackboardReader{
		client: httpclient.New(apiKey),
	}
}

func (r *BlackboardReader) ReadUnreadMessages() ([]domain.Message, error) {
	url := "https://blackboard.example.com/learn/api/public/v1/messages"
	headers := map[string]string{
		"Accept": "application/json",
	}

	resp, err := r.client.Get(url, headers)
	if err != nil {
		return nil, err
	}

	var result struct {
		Messages []domain.Message `json:"results"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return result.Messages, nil
}
