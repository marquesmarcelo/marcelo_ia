package ports

type DraftWriter interface {
	WriteDraft(content string) error
}
