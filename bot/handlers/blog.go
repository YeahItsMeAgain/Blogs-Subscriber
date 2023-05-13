package handlers

import (
	"blogs_subscriber/bot/handlers/utils"
	"blogs_subscriber/db"
	"blogs_subscriber/db/models"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/thoas/go-funk"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm/clause"
)

func HandleListBlogs(ctx telebot.Context) error {
	user := utils.GetCurrentUser(ctx)
	var blogs []string
	for _, blog := range user.Blogs {
		blogs = append(blogs, blog.Url)
	}
	return ctx.Send(strings.Join(blogs, "\n"))
}

func HandleSubscribe(ctx telebot.Context) error {
	return ctx.Send("Send the blog url")
}

func HandleUnsubscribe(ctx telebot.Context) error {
	return ctx.Send("Send the blog url")
}

func HandleBlogRequest(ctx telebot.Context) error {
	blogsUrls := ctx.Text()
	for _, blogUrl := range strings.Split(blogsUrls, "\n") {
		_, err := url.ParseRequestURI(blogUrl)
		if err != nil {
			ctx.Send(fmt.Sprintf("❎ %s is not a valid url", blogUrl))
			continue
		}

		user := utils.GetCurrentUser(ctx)
		blog := models.Blog{Url: blogUrl}
		db.DB.Preload(clause.Associations).FirstOrCreate(&blog, "url = ?", blogUrl)
		if !funk.Contains(blog.Subscribers, func(subscriber *models.User) bool {
			return subscriber.ID == user.ID
		}) {
			db.DB.Model(&blog).Association("Subscribers").Append(&user)
			log.Printf("[*] %d: %s Subscribed to: %s.", ctx.Sender().ID, ctx.Sender().Username, ctx.Text())
			ctx.Send(fmt.Sprintf("✅ You have successfully subscribed to: %s", blogUrl))
			continue
		}

		log.Printf("[*] %d: %s Unsubscribed from: %s.", ctx.Sender().ID, ctx.Sender().Username, ctx.Text())
		db.DB.Model(&blog).Association("Subscribers").Delete(&user)
		ctx.Send(fmt.Sprintf("✅ You have successfully unsubscribed from: %s", blogUrl))
	}
	return nil
}
