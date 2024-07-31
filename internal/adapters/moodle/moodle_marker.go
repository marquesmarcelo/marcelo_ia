package moodle

import (
	"fmt"
	"net/http"
	"net/url"
)

type MoodleMarker struct {
	token    string
	endpoint string
}

func NewMoodleMarker() *MoodleMarker {
	token := "YOUR_MOODLE_TOKEN"
	endpoint := "https://yourmoodlesite.com/webservice/rest/server.php"
	return &MoodleMarker{token: token, endpoint: endpoint}
}

func (m *MoodleMarker) MarkMessageAsRead(messageID string) error {
	data := url.Values{}
	data.Set("moodlewsrestformat", "json")
	data.Set("wsfunction", "core_message_mark_message_read")
	data.Set("wstoken", m.token)
	data.Set("messageid", messageID)

	resp, err := http.PostForm(m.endpoint, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %v", resp.Status)
	}
	return nil
}
