package bot

import (
	"blogs_subscriber/config"
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

func Run() {
	log.Printf("[*] Creating the bot.")
	bot := createBot()

	log.Printf("[*] Starting the bot.")
	bot.Start()
}

func createBot() *telebot.Bot {
	pref := telebot.Settings{
		Token:  config.Config.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	initHandlers(bot)
	return bot
}

func initHandlers(bot *telebot.Bot) {
	bot.Handle("/hello", func(c telebot.Context) error {
		return c.Send("Hello!")
	})
}
