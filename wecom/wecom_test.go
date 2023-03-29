package wecom_test

import (
	"os"
	"testing"

	"github.com/rayepeng/wecombot/wecom"
)

// TestMain() 函数是一个特殊的函数，它允许在测试之前和之后进行一些额外的操作。它被用于在测试运行之前或之后执行初始化或清理操作。
func TestMain(m *testing.M) {
	os.Setenv("WECOM_LOG_MODE", "stdout")
	os.Exit(m.Run())
}

func TestSend(t *testing.T) {
	// 这里使用你的企业微信机器人Webhook URL替换
	webhookURL := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"

	client := wecom.NewClient(webhookURL)

	t.Run("TextMessage", func(t *testing.T) {
		msg := &wecom.TextMessage{
			BaseMessage: wecom.BaseMessage{MsgType: "text"},
			Text:        wecom.Text{Content: "This is a test message"},
		}

		if err := client.Send(msg); err != nil {
			t.Errorf("failed to send text message: %v", err)
		}
	})

	t.Run("MarkdownMessage", func(t *testing.T) {
		msg := &wecom.MarkdownMessage{
			BaseMessage: wecom.BaseMessage{MsgType: "markdown"},
			Markdown:    wecom.Markdown{Content: "### This is a markdown message\n> Hello, world!"},
		}

		if err := client.Send(msg); err != nil {
			t.Errorf("failed to send markdown message: %v", err)
		}
	})

	t.Run("NewsMessage", func(t *testing.T) {
		msg := &wecom.NewsMessage{
			BaseMessage: wecom.BaseMessage{MsgType: "news"},
			News: wecom.News{
				NewsArticle: []*wecom.NewsArticle{
					{
						Title:       "Hello",
						Description: "Hello, world!",
						URL:         "https://www.example.com",
						PicURL:      "https://www.example.com/image.jpg",
					},
				},
			},
		}

		if err := client.Send(msg); err != nil {
			t.Errorf("failed to send news message: %v", err)
		}
	})
}
