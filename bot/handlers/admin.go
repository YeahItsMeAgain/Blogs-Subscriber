package handlers

import (
	"blogs_subscriber/db"
	"blogs_subscriber/db/utils"

	"gopkg.in/telebot.v3"
)

func HandleUsers(ctx telebot.Context) error {
	var users []db.User
	db.DB.Find(&users)
	return ctx.Send(utils.StructsToString(users))
}
