package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TgId      int64 `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Username  string
	Blogs     []Blog `gorm:"many2many:user_blogs;"`
}

type Blog struct {
	gorm.Model
	Url         string  `gorm:"uniqueIndex"`
	Subscribers []*User `gorm:"many2many:user_blogs;"`
	HtmlLen     int     `gorm:"default:0"`
}
