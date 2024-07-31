package memory

import (
	"github.com/marquesmarcelo/marcelo_ia/internal/domain"
)

type MemoryWriter struct {
	messages *[]domain.Message
}

func NewMemoryWriter(messages *[]domain.Message) *MemoryWriter {
	return &MemoryWriter{
		messages: messages,
	}
}

func (m *MemoryWriter) WriteDraft(content string) error {
	*m.messages = append(*m.messages, domain.Message{ID: "3", Subject: "Draft", Body: content, Read: false})
	return nil
}
