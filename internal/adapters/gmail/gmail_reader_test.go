package gmail

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/marquesmarcelo/marcelo_ia/internal/domain"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func TestReadUnreadMessages(t *testing.T) {
	// Mock Gmail API server
	handler := http.NewServeMux()
	handler.HandleFunc("/gmail/v1/users/me/messages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"messages": [{"id": "12345", "threadId": "67890"}]}`))
	})
	handler.HandleFunc("/gmail/v1/users/me/messages/12345", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
            "id": "12345",
            "threadId": "67890",
            "snippet": "Test snippet",
            "payload": {
                "headers": [{"name": "Subject", "value": "Test Subject"}]
            }
        }`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	// Mock Gmail service with the test server URL
	ctx := context.Background()
	srv, err := gmail.NewService(ctx, option.WithEndpoint(server.URL))
	if err != nil {
		t.Fatalf("Unable to create Gmail service: %v", err)
	}

	reader := &GmailReader{
		service: srv,
	}

	got, err := reader.ReadUnreadMessages()
	if err != nil {
		t.Fatalf("ReadUnreadMessages() error = %v", err)
	}

	want := []domain.Message{
		{
			ID:      "12345",
			Subject: "Test Subject",
			Body:    "Test snippet",
		},
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("ReadUnreadMessages() mismatch (-want +got):\n%s", diff)
	}
}
