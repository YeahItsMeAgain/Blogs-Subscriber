package utils

import (
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"fmt"
	"reflect"
	"strconv"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm/clause"
)

func StructsToString[E any](elements []E) string {
	if len(elements) == 0 {
		return "The list is empty."
	}

	var res string
	for _, element := range elements {
		val := reflect.ValueOf(element)
		res += fmt.Sprintf(
			"----------\n%s----------\n", structToString(val),
		)
	}
	return res
}

func structToString(val reflect.Value) string {
	var res string
	for i := 0; i < val.NumField(); i++ {
		valType := val.Type().Field(i)
		if valType.Tag.Get("hidden") != "" {
			continue
		}

		if strVal := valToString(val.Field(i)); strVal != "" {
			res += fmt.Sprintf("%s: %s\n", valType.Name, strVal)
		}
	}
	return res
}

func valToString(val reflect.Value) string {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.String:
		return val.String()
	default:
		return ""
	}
}

func GetCurrentUser(ctx telebot.Context) models.User {
	user := models.User{TgId: ctx.Sender().ID}
	db.DB.Preload(clause.Associations).First(&user, "tg_id = ?", ctx.Sender().ID)
	return user
}
