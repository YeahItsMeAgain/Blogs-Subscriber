package handlers

import (
	"blogs_subscriber/config"
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"blogs_subscriber/db/utils"
	"blogs_subscriber/timer"

	"gopkg.in/telebot.v3"
)

func HandleAdminUsers(ctx telebot.Context) error {
	var users []models.User
	db.DB.Find(&users)
	for _, usersMessage := range utils.StructsToString(users, config.Config.TGMaxMessageLength) {
		err := ctx.Send(usersMessage)
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleAdminBlogs(ctx telebot.Context) error {
	var blogs []models.Blog
	db.DB.Find(&blogs)
	for _, blogsMessage := range utils.StructsToString(blogs, config.Config.TGMaxMessageLength) {
		err := ctx.Send(blogsMessage)
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleAdminUpdate(ctx telebot.Context) error {
	ctx.Send("Updating...")
	timer.UpdateSubscribers(ctx.Bot())
	return ctx.Send("Done.")
}
