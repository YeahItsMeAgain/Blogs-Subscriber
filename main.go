package main

import (
	"blogs_subscriber/bot"
	"blogs_subscriber/config"
	"blogs_subscriber/db"
	"blogs_subscriber/timer"
)

func main() {
	config.Init()
	db.Init()

	tgBot := bot.Init()
	go timer.ScheduleUpdates(tgBot)

	timer.UpdateSubscribers(tgBot)
	bot.Run(tgBot)
}
