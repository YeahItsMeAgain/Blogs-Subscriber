package timer

import (
	"blogs_subscriber/config"
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"blogs_subscriber/timer/utils"
	"fmt"
	"log"
	"math"
	"net/url"
	"time"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm/clause"
)

var timer *time.Timer

func updateTimer() {
	nextTick := time.Date(
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		config.Config.UpdateHourOfDay, 0, 0, 0, time.Local,
	)

	if nextTick.Before(time.Now()) {
		nextTick = nextTick.Add(time.Duration(config.Config.UpdateIntervalHrs) * time.Hour)
	}
	diff := nextTick.Sub(time.Now())
	if timer == nil {
		timer = time.NewTimer(diff)
	} else {
		timer.Reset(diff)
	}
}

func ScheduleUpdates(bot *telebot.Bot) {
	for {
		updateTimer()
		<-timer.C
		UpdateSubscribers(bot)
	}
}

func UpdateSubscribers(bot *telebot.Bot) {
	var blogs []models.Blog
	db.DB.Preload(clause.Associations).Find(&blogs)

	log.Printf("[*] Updating subscribers for new updates.")
	for _, blog := range blogs {
		blogUrl, _ := url.Parse(blog.Url)
		if len(blog.Subscribers) == 0 {
			log.Printf("[*] Deleting %s", blog.Url)
			db.DB.Delete(&blog)
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
					fmt.Sprintf("%s (%s://%s) was updated! (%d)", blog.Url, blogUrl.Scheme, blogUrl.Hostname(), newLength-blog.HtmlLen),
					telebot.NoPreview,
				)
			}
		}
		blog.HtmlLen = newLength
		db.DB.Save(&blog)
	}
	log.Printf("[*] Finished updates.")
}
