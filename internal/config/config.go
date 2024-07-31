package config

import (
	"log"
	"os"
)

type Config struct {
	GmailAPIKey      string
	OutlookClientID  string
	OutlookSecret    string
	OutlookTenantID  string
	MoodleAPIKey     string
	WhatsAppAPIKey   string
	BlackboardAPIKey string
}

func LoadConfig() *Config {
	config := &Config{
		GmailAPIKey:      getEnv("GMAIL_API_KEY", ""),
		OutlookClientID:  getEnv("OUTLOOK_CLIENT_ID", ""),
		OutlookSecret:    getEnv("OUTLOOK_SECRET", ""),
		OutlookTenantID:  getEnv("OUTLOOK_TENANT_ID", ""),
		MoodleAPIKey:     getEnv("MOODLE_API_KEY", ""),
		WhatsAppAPIKey:   getEnv("WHATSAPP_API_KEY", ""),
		BlackboardAPIKey: getEnv("BLACKBOARD_API_KEY", ""),
	}

	// Log a fatal error if any required environment variables are not set
	if config.GmailAPIKey == "" ||
		config.OutlookClientID == "" ||
		config.OutlookSecret == "" ||
		config.OutlookTenantID == "" ||
		config.MoodleAPIKey == "" ||
		config.WhatsAppAPIKey == "" ||
		config.BlackboardAPIKey == "" {
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
