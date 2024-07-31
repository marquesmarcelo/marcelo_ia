package usecases

import (
	"github.com/marquesmarcelo/marcelo_ia/internal/domain"
)

type MessageReader interface {
	ReadUnreadMessages() ([]domain.Message, error)
}

func ReadUnreadMessages(reader MessageReader) ([]domain.Message, error) {
	return reader.ReadUnreadMessages()
}
