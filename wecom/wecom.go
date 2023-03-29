package wecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"

type Client struct {
	webhookURL string
	httpClient *http.Client
}

func NewClient(accessToken string) *Client {
	return &Client{
		webhookURL: fmt.Sprintf(url, accessToken),
		httpClient: &http.Client{},
	}
}

func (client *Client) Send(message Message) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = client.postMessage(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) postMessage(jsonData []byte) error {
	req, err := http.NewRequest("POST", client.webhookURL, bytes.NewReader(jsonData))
	if err != nil {
		Logger.WithError(err).Errorf("failed to create request: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		Logger.WithError(err).Errorf("failed to send message: %v", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Logger.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		Logger.WithError(err).Errorf("unexpected status code: %d", resp.StatusCode)
		return err
	}

	return nil
}

// NewTextMessage 创建一个新的 TextMessage 实例
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		BaseMessage: BaseMessage{MsgType: "text"},
		Text:        Text{Content: content},
	}
}

// NewMarkdownMessage 创建一个新的 MarkdownMessage 实例
func NewMarkdownMessage(content string) *MarkdownMessage {
	return &MarkdownMessage{
		BaseMessage: BaseMessage{MsgType: "markdown"},
		Markdown:    Markdown{Content: content},
	}
}

// NewNewsMessage 创建一个新的 NewsMessage 实例
func NewNewsMessage(news News) *NewsMessage {
	return &NewsMessage{
		BaseMessage: BaseMessage{MsgType: "news"},
		News:        news,
	}
}
