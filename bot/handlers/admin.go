package handlers

import (
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"blogs_subscriber/db/utils"
	"blogs_subscriber/timer"

	"gopkg.in/telebot.v3"
)

func HandleAdminUsers(ctx telebot.Context) error {
	var users []models.User
	db.DB.Find(&users)
	return ctx.Send(utils.StructsToString(users))
}

func HandleAdminBlogs(ctx telebot.Context) error {
	var blogs []models.Blog
	db.DB.Find(&blogs)
	return ctx.Send(utils.StructsToString(blogs))
}

func HandleAdminUpdate(ctx telebot.Context) error {
	ctx.Send("Updating...")
	timer.UpdateSubscribers(ctx.Bot())
	return ctx.Send("Done.")
}
