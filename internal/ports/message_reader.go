package ports

import "github.com/marquesmarcelo/marcelo_ia/domain"

type MessageReader interface {
	ReadUnreadMessages() ([]domain.Message, error)
}
