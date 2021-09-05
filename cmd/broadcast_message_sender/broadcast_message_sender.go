package main

import (
	"fmt"

	"github.com/kobaryubi/vc-mop/internal/line"
)

func main() {
	message := make(map[string]string, 2)
	message["type"] = "text"
	message["text"] = "Hello"

	params := line.PostBroadcastMessageParams{
		Messages:             make([]map[string]string, 1),
		NotificationDisabled: true,
	}
	params.Messages = append(params.Messages, message)

	lineApiClient := line.NewClient()
	fmt.Println(lineApiClient)
}
