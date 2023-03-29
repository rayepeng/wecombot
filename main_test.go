package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	app := createApp()

	input := []string{
		"wecom",
		"--access_token", "your_access_token",
		"--text", "Hello, WeCom!",
	}

	err := app.Run(input)
	assert.NoError(t, err)

	// 如果你需要检查输出，你可以设置 app.Writer 为一个 bytes.Buffer 实例，然后检查它的内容。
	outputBuffer := &bytes.Buffer{}
	app.Writer = outputBuffer

	// 重新运行应用程序以捕获输出。
	err = app.Run(input)
	assert.NoError(t, err)

	output := outputBuffer.String()
	assert.Contains(t, output, "Message sent successfully")
}
