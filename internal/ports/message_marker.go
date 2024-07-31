package ports

type MessageMarker interface {
    MarkMessageAsRead(messageID string) error
}
