package handlers

import (
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"blogs_subscriber/db/utils"

	"gopkg.in/telebot.v3"
)

func HandleList(ctx telebot.Context) error {
	var blogs []models.BaseBlog
	db.DB.Model(&models.Blog{}).Find(&blogs)
	return ctx.Send(utils.StructsToString(blogs))
}
