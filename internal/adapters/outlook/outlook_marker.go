package outlook

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/oauth2/clientcredentials"
)

type OutlookMarker struct {
	client *msgraphsdk.GraphServiceClient
}

func NewOutlookMarker() *OutlookMarker {
	clientID := "YOUR_CLIENT_ID"
	clientSecret := "YOUR_CLIENT_SECRET"
	tenantID := "YOUR_TENANT_ID"

	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://login.microsoftonline.com/" + tenantID + "/oauth2/v2.0/token",
		Scopes:       []string{"https://graph.microsoft.com/.default"},
	}

	ctx := context.Background()
	httpClient := config.Client(ctx)
	client := msgraphsdk.NewGraphServiceClient(httpClient)

	return &OutlookMarker{client: client}
}

func (m *OutlookMarker) MarkMessageAsRead(messageID string) error {
	msgID := models.NewMessage()
	msgID.SetIsRead(true)

	request := m.client.Me().MessagesById(messageID).Request()
	err := request.Update(context.Background(), msgID)
	if err != nil {
		return err
	}
	return nil
}
