package line

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kobaryubi/vc-mop/internal/sender"
)

const messagingAPIBaseURL = "https://api.line.me/v2/bot/message"

// APIClient API用のクライアント
type APIClient struct{}

// NewAPIClient 初期化用の関数
func NewAPIClient() *APIClient {
	apiClient := &APIClient{}
	return apiClient
}

// BroadcastMessageParams POSTリクエスト用のパラメータ
type BroadcastMessageParams struct {
	Messages             []Message `json:"messages"`
	NotificationDisabled bool      `json:"notificationDisabled"`
}

// Message Messageの定義
type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// NewMessage Message初期化関数
func NewMessage(text string) Message {
	message := Message{"text", text}
	return message
}

// SendBroadcastMessage ブロードキャストメッセージ送信用の関数
func (apiClient *APIClient) SendBroadcastMessage(params *BroadcastMessageParams) (body []byte, err error) {
	err = godotenv.Load()
	if err != nil {
		// TODO: ログ出力log.Fatal
		return
	}

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer {%s}", os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")),
	}

	httpSender := sender.NewHTTPSender(messagingAPIBaseURL+"/broadcast", "POST")
	paramsJSON, _ := json.Marshal(params)

	byteArray, err := httpSender.Send(paramsJSON, header)
	return byteArray, err
}
