package usecases

type DraftWriter interface {
	WriteDraft(content string) error
}

func WriteDraft(writer DraftWriter, content string) error {
	return writer.WriteDraft(content)
}
