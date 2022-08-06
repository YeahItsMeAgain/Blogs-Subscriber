package utils

import (
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm/clause"
)

func GetCurrentUser(ctx telebot.Context) models.User {
	user := models.User{TgId: ctx.Sender().ID}
	db.DB.Preload(clause.Associations).First(&user, "tg_id = ?", ctx.Sender().ID)
	return user
}
