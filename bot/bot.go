package bot

import (
	"blogs_subscriber/bot/handlers"
	"blogs_subscriber/bot/menu"
	"blogs_subscriber/bot/utils"
	"blogs_subscriber/config"
	"log"

	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func Run() {
	log.Printf("[*] Creating the bot.")
	bot := utils.CreateBot(config.Config.BotToken)

	log.Printf("[*] Creating handlers for the bot.")
	initHandlers(bot)

	log.Printf("[*] Starting the bot.")
	bot.Start()
}

func initHandlers(bot *telebot.Bot) {
	bot.Handle("/start", handlers.HandleStart)
	bot.Handle(&menu.BtnList, handlers.HandleListBlogs)
	bot.Handle(&menu.BtnSubscribe, handlers.HandleSubscribe)
	bot.Handle(&menu.BtnUnsubscribe, handlers.HandleUnsubscribe)
	bot.Handle(telebot.OnText, handlers.HandleBlogRequest)

	adminOnly := bot.Group()
	adminOnly.Use(middleware.Whitelist(config.Config.AdminIds...))
	adminOnly.Handle(&menu.AdminBtnUsers, handlers.HandleAdminUsers)
	adminOnly.Handle(&menu.AdminBtnBlogs, handlers.HandleAdminBlogs)
}
