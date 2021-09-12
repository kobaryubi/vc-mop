package main

import (
	"fmt"
	"os"

	"github.com/kobaryubi/vc-mop/internal/sender"
)

func main() {
	msg := "Hello"
	sender := sender.NewHTTPSender("http://localhost:8080")
	data, err := sender.Send([]byte(msg))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
