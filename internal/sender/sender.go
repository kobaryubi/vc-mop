package sender

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPSender HTTPリクエスト送信用の定義
type HTTPSender struct {
	host   string
	method string
	client *http.Client
}

// NewHTTPSender HTTPプロトコルの通信クライアントを作成する関数
func NewHTTPSender(host string, method string) *HTTPSender {
	httpSender := &HTTPSender{host, method, &http.Client{}}
	return httpSender
}

// Send 送信用の関数
func (httpSender *HTTPSender) Send(data []byte, header map[string]string) (body []byte, err error) {
	req, err := http.NewRequest(httpSender.method, httpSender.host, bytes.NewBuffer(data))
	for key, value := range header {
		req.Header.Set(key, value)
	}

	if err != nil {
		return nil, err
	}

	resp, err := httpSender.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
