package main

import (
	"fmt"

	"github.com/marquesmarcelo/marcelo_ia/internal/config"
	"github.com/marquesmarcelo/marcelo_ia/internal/usecases"
)

func main() {
	cfg := config.LoadConfig()
	app := config.Setup(cfg)

	readers := []usecases.MessageReader{app.GmailReader, app.OutlookReader, app.MoodleReader, app.WhatsAppReader, app.BlackboardReader}

	for _, reader := range readers {
		messages, err := usecases.ReadUnreadMessages(reader)
		if err != nil {
			fmt.Printf("Error reading messages: %v\n", err)
			continue
		}
		fmt.Printf("Unread messages from %T: %v\n", reader, messages)

		for _, msg := range messages {
			err := usecases.WriteDraft(app.GmailWriter, "This is a test draft message")
			if err != nil {
				fmt.Printf("Error writing draft: %v\n", err)
				continue
			}
			fmt.Println("Draft written successfully.")

			err = usecases.MarkMessageAsRead(app.GmailMarker, msg.ID)
			if err != nil {
				fmt.Printf("Error marking message as read: %v\n", err)
			} else {
				fmt.Println("Message marked as read successfully.")
			}
		}
	}
}
