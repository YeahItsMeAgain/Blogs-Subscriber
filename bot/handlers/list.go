package handlers

import "gopkg.in/telebot.v3"

func HandleList(ctx telebot.Context) error {
	return ctx.Send("Listing")
}
