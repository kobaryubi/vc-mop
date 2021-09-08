package line

const messagingAPIBaseURL = "https://api.line.me/v2/bot/message"

// APIClient API用のクライアント
type APIClient struct{}

// NewAPIClient 初期化用の関数
func NewAPIClient() *APIClient {
	apiClient := &APIClient{}
	return apiClient
}

// PostBroadcastMessageParams POSTリクエスト用のパラメータ
type PostBroadcastMessageParams struct {
	Messages             []map[string]string
	NotificationDisabled bool
}
