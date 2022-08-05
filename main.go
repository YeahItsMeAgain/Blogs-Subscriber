package main

import (
	"blogs_subscriber/bot"
	"blogs_subscriber/config"
)

func main() {
	config.Init()
	bot.Run()
}
