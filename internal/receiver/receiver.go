package receiver

import (
	"fmt"
	"net/http"
)

type HTTPReceiver struct {
	pattern string
	port    int
}

func NewHTTPReceiver(pattern string, port int) *HTTPReceiver {
	httpReceiver := &HTTPReceiver{pattern, port}
	return httpReceiver
}

func (httpReceiver *HTTPReceiver) Receive() {
	http.HandleFunc(httpReceiver.pattern, httpReceiver.handler)
	http.ListenAndServe(fmt.Sprintf(":%d", httpReceiver.port), nil)
}

func (httpReceiver *HTTPReceiver) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}
