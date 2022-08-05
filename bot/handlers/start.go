package handlers

import (
	"blogs_subscriber/bot/menu"

	"gopkg.in/telebot.v3"
)

func HandleStart(ctx telebot.Context) error {
	return ctx.Send("Menu", menu.Menu)
}
