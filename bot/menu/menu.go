package menu

import (
	"blogs_subscriber/bot/utils"

	"gopkg.in/telebot.v3"
)

var (
	BtnList       = telebot.Btn{Text: "📚 List"}
	AdminBtnUsers = telebot.Btn{Text: "👤 Users"}
	AdminBtnBlogs = telebot.Btn{Text: "📝 Blogs"}
	menuButtons   = telebot.Row{BtnList}
	Menu          = &telebot.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: utils.CreateReplyMarkup(
			menuButtons,
		),
	}
	AdminMenu = &telebot.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: utils.CreateReplyMarkup(
			menuButtons,
			telebot.Row{AdminBtnUsers, AdminBtnBlogs},
		),
	}
)
