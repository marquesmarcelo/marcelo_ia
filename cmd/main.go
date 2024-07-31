package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/marquesmarcelo/marcelo_ia/internal/config"
	"github.com/marquesmarcelo/marcelo_ia/internal/usecases"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := config.LoadConfig()
	app := config.Setup(cfg)

	// Usar o adaptador em memória para ler mensagens
	messages, err := usecases.ReadUnreadMessages(app.MemoryReader)
	if err != nil {
		log.Fatalf("Error reading messages: %v", err)
	}
	fmt.Printf("Unread messages: %v\n", messages)

	// Escrever um rascunho usando o adaptador em memória
	err = usecases.WriteDraft(app.MemoryWriter, "This is a draft message")
	if err != nil {
		log.Fatalf("Error writing draft: %v", err)
	}
	fmt.Println("Draft written successfully.")

	// Marcar a primeira mensagem como lida
	if len(messages) > 0 {
		err = usecases.MarkMessageAsRead(app.MemoryMarker, messages[0].ID)
		if err != nil {
			log.Fatalf("Error marking message as read: %v", err)
		}
		fmt.Println("Message marked as read successfully.")
	}

	// Listar as mensagens em memória para verificar o rascunho escrito e a marcação como lida
	finalMessages, err := usecases.ReadUnreadMessages(app.MemoryReader)
	if err != nil {
		log.Fatalf("Error reading messages: %v", err)
	}
	fmt.Printf("Final messages: %v\n", finalMessages)
}
