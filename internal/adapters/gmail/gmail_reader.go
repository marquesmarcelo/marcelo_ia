package gmail

import (
	"context"
	"log"

	"github.com/marquesmarcelo/marcelo_ia/domain"
	"github.com/marquesmarcelo/marcelo_iapkg/httpclient"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailReader struct {
	client *httpclient.HTTPClient
}

func NewGmailReader(apiKey string) *GmailReader {
	return &GmailReader{
		client: httpclient.New(apiKey),
	}
}

func (r *GmailReader) ReadUnreadMessages() ([]domain.Message, error) {
	ctx := context.Background()
	srv, err := gmail.NewService(ctx, option.WithAPIKey(r.client.APIKey))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	res, err := srv.Users.Messages.List(user).LabelIds("INBOX").Q("is:unread").Do()
	if err != nil {
		return nil, err
	}

	var messages []domain.Message
	for _, m := range res.Messages {
		msg, err := srv.Users.Messages.Get(user, m.Id).Do()
		if err != nil {
			return nil, err
		}
		messages = append(messages, domain.Message{
			ID:      msg.Id,
			Subject: msg.Payload.Headers[0].Value,
			Body:    msg.Snippet,
		})
	}

	return messages, nil
}
