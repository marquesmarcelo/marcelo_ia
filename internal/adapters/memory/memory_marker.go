package memory

import (
	"github.com/marquesmarcelo/marcelo_ia/internal/domain"
)

type MemoryMarker struct {
	messages *[]domain.Message
}

func NewMemoryMarker(messages *[]domain.Message) *MemoryMarker {
	return &MemoryMarker{
		messages: messages,
	}
}

func (m *MemoryMarker) MarkMessageAsRead(messageID string) error {
	for i, msg := range *m.messages {
		if msg.ID == messageID {
			(*m.messages)[i].Read = true
		}
	}
	return nil
}
