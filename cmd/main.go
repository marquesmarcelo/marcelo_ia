package main

import (
    "fmt"
    "github.com/marquesmarcelo/marcelo_ia/adapters/gmail"
    "github.com/marquesmarcelo/marcelo_ia/adapters/outlook"
    "github.com/marquesmarcelo/marcelo_ia/adapters/moodle"
    "github.com/marquesmarcelo/marcelo_ia/adapters/whatsapp"
    "github.com/marquesmarcelo/marcelo_ia/adapters/blackboard"
    "github.com/marquesmarcelo/marcelo_ia/config"
    "github.com/marquesmarcelo/marcelo_ia/usecases"
)

func main() {
    cfg := config.LoadConfig()

    gmailReader := gmail.NewGmailReader(cfg.GmailAPIKey)
    outlookReader := outlook.NewOutlookReader(cfg.OutlookClientID, cfg.OutlookSecret, cfg.OutlookTenantID)
    moodleReader := moodle.NewMoodleReader(cfg.MoodleAPIKey)
    whatsappReader := whatsapp.NewWhatsAppReader(cfg.WhatsAppAPIKey)
    blackboardReader := blackboard.NewBlackboardReader(cfg.BlackboardAPIKey)

    readers := []usecases.MessageReader{gmailReader, outlookReader, moodleReader, whatsappReader, blackboardReader}

    for _, reader := range readers {
        messages, err := usecases.ReadUnreadMessages(reader)
        if err != nil {
            fmt.Printf("Error reading messages: %v\n", err)
            continue
        }
        fmt.Printf("Unread messages from %T: %v\n", reader, messages)

        for _, msg := range messages {
            whatsappWriter := whatsapp.NewWhatsAppWriter(cfg.WhatsApp
