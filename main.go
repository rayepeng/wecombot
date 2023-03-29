package main

import (
	"fmt"
	"os"

	"github.com/rayepeng/simplecmdtool/tools"
	"github.com/rayepeng/wecombot/wecom"
	"github.com/urfave/cli/v2"
)

func PushWeComMsg(ctx *cli.Context) error {
	accessToken := ctx.String("a")
	text := ctx.String("m")

	msg := wecom.NewTextMessage(text)
	client := wecom.NewClient(accessToken)
	err := client.Send(msg)
	if err != nil {
		return err
	}

	fmt.Println("Message sent successfully")
	return nil
}

func main() {
	tool := tools.NewTool([]*tools.FunctionConfig{
		{
			Name:        "PushWeComMsg",
			Function:    PushWeComMsg,
			OptionFlags: "a:m:",
			Description: "-a <access_token> -m <msg>",
		},
	})
	err := tool.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
