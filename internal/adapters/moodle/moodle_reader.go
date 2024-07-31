package moodle

import (
	"github.com/marquesmarcelo/marcelo_ia/domain"
)

type MoodleReader struct{}

func NewMoodleReader() *MoodleReader {
	return &MoodleReader{}
}

func (r *MoodleReader) ReadUnreadMessages() ([]domain.Message, error) {
	// Implementar a lógica de leitura de mensagens não lidas do Moodle
	return []domain.Message{
		{ID: "3", Subject: "Moodle Subject 1", Body: "Moodle Body 1"},
	}, nil
}
