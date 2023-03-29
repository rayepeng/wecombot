package main

import (
	"os"
	"testing"

	"github.com/rayepeng/simplecmdtool/tools"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {

	tool := tools.NewTool([]*tools.FunctionConfig{
		{
			Name:        "PushWeComMsg",
			Function:    PushWeComMsg,
			OptionFlags: "a:m:",
			Description: "-a <access_token> -m <msg>",
		},
	})
	access_token := os.Getenv("WECOM_ACCESS_TOKEN")
	if access_token == "" {
		assert.Error(t, nil)
	}
	args := []string{"appName", "PushWeComMsg", "-a", access_token, "-m", "hello bot"}
	err := tool.Run(args)
	assert.NoError(t, err)
}
