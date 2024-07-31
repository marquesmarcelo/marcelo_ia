package config

import (
	"log"
	"os"
)

type Config struct {
	GmailAPIKey      string
	GmailAPIURL      string
	OutlookClientID  string
	OutlookSecret    string
	OutlookTenantID  string
	OutlookAPIURL    string
	MoodleAPIKey     string
	MoodleAPIURL     string
	WhatsAppAPIKey   string
	WhatsAppAPIURL   string
	BlackboardAPIKey string
	BlackboardAPIURL string
}

func LoadConfig() *Config {
	config := &Config{
		GmailAPIKey:      getEnv("GMAIL_API_KEY", ""),
		GmailAPIURL:      getEnv("GMAIL_API_URL", "https://gmail.googleapis.com/gmail/v1"),
		OutlookClientID:  getEnv("OUTLOOK_CLIENT_ID", ""),
		OutlookSecret:    getEnv("OUTLOOK_SECRET", ""),
		OutlookTenantID:  getEnv("OUTLOOK_TENANT_ID", ""),
		OutlookAPIURL:    getEnv("OUTLOOK_API_URL", "https://graph.microsoft.com/v1.0"),
		MoodleAPIKey:     getEnv("MOODLE_API_KEY", ""),
		MoodleAPIURL:     getEnv("MOODLE_API_URL", "https://yourmoodlesite.com/webservice/rest/server.php"),
		WhatsAppAPIKey:   getEnv("WHATSAPP_API_KEY", ""),
		WhatsAppAPIURL:   getEnv("WHATSAPP_API_URL", "https://api.whatsapp.com/v1/messages"),
		BlackboardAPIKey: getEnv("BLACKBOARD_API_KEY", ""),
		BlackboardAPIURL: getEnv("BLACKBOARD_API_URL", "https://blackboard.example.com/learn/api/public/v1"),
	}

	// Log a fatal error if any required environment variables are not set
	if config.GmailAPIKey == "" || config.GmailAPIURL == "" ||
		config.OutlookClientID == "" || config.OutlookSecret == "" || config.OutlookTenantID == "" || config.OutlookAPIURL == "" ||
		config.MoodleAPIKey == "" || config.MoodleAPIURL == "" ||
		config.WhatsAppAPIKey == "" || config.WhatsAppAPIURL == "" ||
		config.BlackboardAPIKey == "" || config.BlackboardAPIURL == "" {
		log.Fatal("Missing required environment variables")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
