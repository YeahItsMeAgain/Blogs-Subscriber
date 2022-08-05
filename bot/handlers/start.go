package handlers

import (
	"blogs_subscriber/bot/menu"
	"blogs_subscriber/config"
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"

	"golang.org/x/exp/slices"
	"gopkg.in/telebot.v3"
)

func HandleStart(ctx telebot.Context) error {
	db.DB.FirstOrCreate(&models.User{
		BaseUser: models.BaseUser{
			TgId:      ctx.Sender().ID,
			FirstName: ctx.Sender().FirstName,
			LastName:  ctx.Sender().LastName,
			Username:  ctx.Sender().Username,
		},
	})

	if slices.Contains(config.Config.AdminIds, ctx.Sender().ID) {
		return ctx.Send("Menu", menu.AdminMenu)
	}
	return ctx.Send("Menu", menu.Menu)
}
