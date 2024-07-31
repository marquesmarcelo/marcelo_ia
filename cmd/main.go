package main

import (
	"fmt"

	"github.com/marquesmarcelo/marcelo_ia/internal/config"
	"github.com/marquesmarcelo/marcelo_ia/internal/usecases"
)

func main() {
	cfg := config.LoadConfig()
	app := config.Setup(cfg)

	err := usecases.WriteDraft(app.GmailWriter, "This is a test draft message")
	if err != nil {
		fmt.Printf("Error writing draft: %v\n", err)
	} else {
		fmt.Println("Draft written successfully.")
	}
}
