package timer

import (
	"blogs_subscriber/config"
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"blogs_subscriber/timer/utils"
	"fmt"
	"log"
	"math"
	"time"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm/clause"
)

func ScheduleUpdates(bot *telebot.Bot) {
	ticker := time.NewTicker(time.Duration(config.Config.UpdateIntervalHrs) * time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				UpdateSubscribers(bot)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func UpdateSubscribers(bot *telebot.Bot) {
	var blogs []models.Blog
	db.DB.Preload(clause.Associations).Find(&blogs)

	log.Printf("[*] Updating subscribers for new updates.")
	for _, blog := range blogs {
		if len(blog.Subscribers) == 0 {
			continue
		}

		newLength, err := utils.GetHtmlLength(blog.Url)
		if err != nil {
			log.Printf("[!] Error updating %s: %s", blog.Url, err)
			continue
		}

		// If the difference is large enough
		if (int)(math.Abs((float64)(newLength-blog.HtmlLen))) > config.Config.AllowedErrorMargin {
			log.Printf("[*] %s was updated from %d to %d", blog.Url, blog.HtmlLen, newLength)
			for _, subscriber := range blog.Subscribers {
				bot.Send(
					&telebot.User{ID: subscriber.TgId},
					fmt.Sprintf("%s was updated!", blog.Url),
					telebot.NoPreview,
				)
			}
		}
		blog.HtmlLen = newLength
		db.DB.Save(&blog)
	}
	log.Printf("[*] Finished updates.")
}
