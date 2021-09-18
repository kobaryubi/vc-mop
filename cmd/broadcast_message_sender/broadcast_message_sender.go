package main

import (
	"fmt"

	"github.com/kobaryubi/vc-mop/internal/line"
)

func main() {
	message := line.NewMessage("Hello")
	messages := make([]line.Message, 0)
	messages = append(messages, message)

	broadcastMessageParams := new(line.BroadcastMessageParams)
	broadcastMessageParams.Messages = messages
	broadcastMessageParams.NotificationDisabled = false

	apiClient := line.NewAPIClient()
	byteArray, _ := apiClient.SendBroadcastMessage(broadcastMessageParams)
	fmt.Println(string(byteArray))
}
