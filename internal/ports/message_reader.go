package ports

import "github.com/marquesmarcelo/marcelo_ia/internal/domain"

type MessageReader interface {
	ReadUnreadMessages() ([]domain.Message, error)
}
