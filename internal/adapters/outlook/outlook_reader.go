package outlook

import (
	"github.com/marquesmarcelo/marcelo_ia/domain"
)

type OutlookReader struct{}

func NewOutlookReader() *OutlookReader {
	return &OutlookReader{}
}

func (r *OutlookReader) ReadUnreadMessages() ([]domain.Message, error) {
	// Implementar a lógica de leitura de mensagens não lidas do Outlook
	return []domain.Message{
		{ID: "2", Subject: "Outlook Subject 1", Body: "Outlook Body 1"},
	}, nil
}
