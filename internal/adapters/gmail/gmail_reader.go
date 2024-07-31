package gmail

import (
	"context"
	"log"

	"github.com/marquesmarcelo/marcelo_ia/internal/domain"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailReader struct {
	service *gmail.Service
}

func NewGmailReader(apiKey string) *GmailReader {
	ctx := context.Background()
	srv, err := gmail.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	return &GmailReader{
		service: srv,
	}
}

func (r *GmailReader) ReadUnreadMessages() ([]domain.Message, error) {
	user := "me"
	res, err := r.service.Users.Messages.List(user).LabelIds("INBOX").Q("is:unread").Do()
	if err != nil {
		return nil, err
	}

	var messages []domain.Message
	for _, m := range res.Messages {
		msg, err := r.service.Users.Messages.Get(user, m.Id).Do()
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
