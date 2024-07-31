package config

import (
	"github.com/marquesmarcelo/marcelo_ia/internal/adapters/gmail"
	"github.com/marquesmarcelo/marcelo_ia/internal/adapters/memory"
	"github.com/marquesmarcelo/marcelo_ia/internal/domain"
	"github.com/marquesmarcelo/marcelo_ia/internal/httpclient"
)

type Application struct {
	GmailReader  *gmail.GmailReader
	GmailWriter  *gmail.GmailWriter
	GmailMarker  *gmail.GmailMarker
	MemoryReader *memory.MemoryReader
	MemoryWriter *memory.MemoryWriter
	MemoryMarker *memory.MemoryMarker
}

func Setup(cfg *Config) *Application {
	httpClient := httpclient.New(cfg.GmailAPIKey)
	messages := []domain.Message{
		{ID: "1", Subject: "Test 1", Body: "This is a test message 1", Read: false},
		{ID: "2", Subject: "Test 2", Body: "This is a test message 2", Read: false},
	}

	return &Application{
		GmailReader:  gmail.NewGmailReader(cfg.GmailAPIKey),
		GmailWriter:  gmail.NewGmailWriter(httpClient, cfg.GmailAPIURL),
		GmailMarker:  gmail.NewGmailMarker(httpClient, cfg.GmailAPIURL),
		MemoryReader: memory.NewMemoryReader(&messages),
		MemoryWriter: memory.NewMemoryWriter(&messages),
		MemoryMarker: memory.NewMemoryMarker(&messages),
	}
}
