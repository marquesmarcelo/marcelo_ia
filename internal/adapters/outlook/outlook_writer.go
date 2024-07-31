package outlook

type OutlookWriter struct{}

func NewOutlookWriter() *OutlookWriter {
	return &OutlookWriter{}
}

func (w *OutlookWriter) WriteDraft(content string) error {
	// Implementar a lógica de escrita de rascunhos no Outlook
	return nil
}
