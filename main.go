package main

import (
	"fmt"
	"os"

	"github.com/rayepeng/wecombot/wecom"
	"github.com/urfave/cli/v2"
)

func createApp() *cli.App {
	app := &cli.App{
		Name:  "wecom",
		Usage: "Send messages to WeCom",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "access_token",
				Usage:    "Access token for WeCom",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "text",
				Usage:    "Text message to send",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			accessToken := c.String("access_token")
			text := c.String("text")

			msg := wecom.NewTextMessage(text)
			client := wecom.NewClient(accessToken)
			err := client.Send(msg)
			if err != nil {
				return err
			}

			fmt.Println("Message sent successfully")
			return nil
		},
	}
	return app
}

func main() {
	app := createApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
