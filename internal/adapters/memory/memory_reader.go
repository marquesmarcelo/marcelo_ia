package memory

import (
	"github.com/marquesmarcelo/marcelo_ia/internal/domain"
)

type MemoryReader struct {
	messages *[]domain.Message
}

func NewMemoryReader(messages *[]domain.Message) *MemoryReader {
	return &MemoryReader{
		messages: messages,
	}
}

func (m *MemoryReader) ReadUnreadMessages() ([]domain.Message, error) {
	var unreadMessages []domain.Message
	for _, msg := range *m.messages {
		if !msg.Read {
			unreadMessages = append(unreadMessages, msg)
		}
	}
	return unreadMessages, nil
}
