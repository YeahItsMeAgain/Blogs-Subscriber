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

func Init() *telebot.Bot {
	log.Printf("[*] Creating the bot.")
	bot := utils.CreateBot(config.Config.BotToken)

	log.Printf("[*] Creating handlers for the bot.")
	initHandlers(bot)
	return bot
}

func Run(bot *telebot.Bot) {
	log.Printf("[*] Starting the bot.")
	bot.Start()
}

func initHandlers(bot *telebot.Bot) {
	onlyUsers := bot.Group()
	onlyUsers.Use(middleware.Whitelist(config.Config.AllowedIds...))
	onlyUsers.Handle("/start", handlers.HandleStart)
	onlyUsers.Handle(&menu.BtnList, handlers.HandleListBlogs)
	onlyUsers.Handle(&menu.BtnSubscribe, handlers.HandleSubscribe)
	onlyUsers.Handle(&menu.BtnUnsubscribe, handlers.HandleUnsubscribe)
	onlyUsers.Handle(telebot.OnText, handlers.HandleBlogRequest)

	adminOnly := bot.Group()
	adminOnly.Use(middleware.Whitelist(config.Config.AdminIds...))
	adminOnly.Handle(&menu.AdminBtnUsers, handlers.HandleAdminUsers)
	adminOnly.Handle(&menu.AdminBtnBlogs, handlers.HandleAdminBlogs)
	adminOnly.Handle(&menu.AdminBtnUpdate, handlers.HandleAdminUpdate)
}
