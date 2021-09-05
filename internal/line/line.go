package line

const messagingAPIBaseURL = "https://api.line.me/v2/bot/message"

type APIClient struct{}

func NewClient() *APIClient {
	apiClient := &APIClient{}
	return apiClient
}

type PostBroadcastMessageParams struct {
	Messages             []map[string]string
	NotificationDisabled bool
}
