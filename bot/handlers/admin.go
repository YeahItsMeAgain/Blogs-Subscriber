package handlers

import (
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"blogs_subscriber/db/utils"

	"gopkg.in/telebot.v3"
)

func HandleUsers(ctx telebot.Context) error {
	var users []models.BaseUser
	db.DB.Model(&models.User{}).Find(&users)
	return ctx.Send(utils.StructsToString(users))
}

func HandleBlogs(ctx telebot.Context) error {
	var blogs []models.BaseBlog
	db.DB.Model(&models.Blog{}).Find(&blogs)
	return ctx.Send(utils.StructsToString(blogs))
}
