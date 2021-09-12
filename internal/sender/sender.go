package sender

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPSender HTTPリクエスト送信用の定義
type HTTPSender struct {
	host   string
	client *http.Client
}

// NewHTTPSender HTTPプロトコルの通信クライアントを作成する関数
func NewHTTPSender(host string) *HTTPSender {
	httpSender := &HTTPSender{host, &http.Client{}}
	return httpSender
}

func (httpSender *HTTPSender) Send(data []byte) (body []byte, err error) {
	req, err := http.NewRequest("GET", httpSender.host, bytes.NewBuffer(data))
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
