package config

import (
	"github.com/marquesmarcelo/marcelo_ia/internal/adapters/gmail"
	"github.com/marquesmarcelo/marcelo_ia/pkg/httpclient"
)

type Application struct {
	GmailReader *gmail.GmailReader
	GmailWriter *gmail.GmailWriter
	GmailMarker *gmail.GmailMarker
}

func Setup(cfg *Config) *Application {
	httpClient := httpclient.New(cfg.GmailAPIKey)

	return &Application{
		GmailReader: gmail.NewGmailReader(cfg.GmailAPIKey),
		GmailWriter: gmail.NewGmailWriter(httpClient, cfg.GmailAPIURL),
		GmailMarker: gmail.NewGmailMarker(httpClient, cfg.GmailAPIURL),
	}
}
