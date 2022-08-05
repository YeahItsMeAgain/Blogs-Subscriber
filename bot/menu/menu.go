package menu

import (
	"blogs_subscriber/bot/utils"

	"gopkg.in/telebot.v3"
)

var (
	BtnList = telebot.Btn{Text: "📚 List"}
	Menu    = &telebot.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: utils.CreateReplyMarkup(
			telebot.Row{BtnList},
		),
	}
)
