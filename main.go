package main

import (
	"blogs_subscriber/bot"
	"blogs_subscriber/config"
	"blogs_subscriber/db"
)

func main() {
	config.Init()
	db.Init()
	bot.Run()
}
