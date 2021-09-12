package receiver

import (
	"fmt"
	"net/http"
)

// HTTPReceiver HTTP通信用のreceiver
type HTTPReceiver struct {
	pattern string
	port    int
}

// NewHTTPReceiver HTTPReceiverの初期化関数
func NewHTTPReceiver(pattern string, port int) *HTTPReceiver {
	httpReceiver := &HTTPReceiver{pattern, port}
	return httpReceiver
}

// Receive Receive関数
func (httpReceiver *HTTPReceiver) Receive() {
	http.HandleFunc(httpReceiver.pattern, httpReceiver.handler)
	http.ListenAndServe(fmt.Sprintf(":%d", httpReceiver.port), nil)
}

// handler handler関数
func (httpReceiver *HTTPReceiver) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}
