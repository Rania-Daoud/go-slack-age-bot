package main

import (
	"context"
	"log"
	"os"

	"github.com/rania-daoud/go-slack-bot/utils"
	"github.com/shomali11/slacker"
)

const (
	SLACK_BOT_TOKEN = "SLACK_BOT_TOKEN"
	SLACK_APP_TOKEN = "SLACK_APP_TOKEN"
)

func main() {
	os.Setenv(SLACK_BOT_TOKEN, "xoxb-4087163464405-4102805721649-BtLjVicwLzrkvLNhhVU7lx9p")
	os.Setenv(SLACK_APP_TOKEN, "xapp-1-A042FFE7LLE-4102770478689-60b46bbc03dac4b5e2d50eec621ce9963b9f2774c47b88b4a53f819612de209c")

	bot := slacker.NewClient(os.Getenv(SLACK_BOT_TOKEN), os.Getenv(SLACK_APP_TOKEN))

	go utils.PrintCommandEvents(bot.CommandEvents())

	bot.Command("My yob is: <year>", utils.SlackerCommandDefintion)

	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(context)

	if err != nil {
		log.Fatal(err)
	}
}
