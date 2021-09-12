package main

import (
	"github.com/kobaryubi/vc-mop/internal/receiver"
)

func main() {
	pattern := "/"
	port := 8080
	httpReceiver := receiver.NewHTTPReceiver(pattern, port)
	httpReceiver.Receive()
}
