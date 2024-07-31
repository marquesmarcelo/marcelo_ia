package moodle

type MoodleWriter struct{}

func NewMoodleWriter() *MoodleWriter {
	return &MoodleWriter{}
}

func (w *MoodleWriter) WriteDraft(content string) error {
	// Implementar a lógica de escrita de rascunhos no Moodle
	return nil
}
