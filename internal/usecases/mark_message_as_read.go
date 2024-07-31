package usecases

type MessageMarker interface {
    MarkMessageAsRead(messageID string) error
}

func MarkMessageAsRead(marker MessageMarker, messageID string) error {
    return marker.MarkMessageAsRead(messageID)
}
