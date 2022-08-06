package handlers

import (
	"blogs_subscriber/bot/menu"
	"blogs_subscriber/config"
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"fmt"
	"log"

	"golang.org/x/exp/slices"
	"gopkg.in/telebot.v3"
)

func HandleStart(ctx telebot.Context) error {
	db.DB.FirstOrCreate(&models.User{
		TgId:      ctx.Sender().ID,
		FirstName: ctx.Sender().FirstName,
		LastName:  ctx.Sender().LastName,
		Username:  ctx.Sender().Username,
	}, "tg_id = ?", ctx.Sender().ID)
	log.Printf("[*] %d : %s Started the bot.", ctx.Sender().ID, ctx.Sender().Username)

	if slices.Contains(config.Config.AdminIds, ctx.Sender().ID) {
		return ctx.Send(fmt.Sprintf("👋 Welcome %s!", ctx.Sender().FirstName), menu.AdminMenu)
	}
	return ctx.Send("👋 Welcome!", menu.Menu)
}
