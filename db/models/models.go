package models

import (
	"gorm.io/gorm"
)

type BaseUser struct {
	gorm.Model
	TgId      int64 `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Username  string
}

type BaseBlog struct {
	gorm.Model
	Url string
}

type User struct {
	BaseUser
	Blogs []Blog `gorm:"many2many:user_blogs;"`
}

type Blog struct {
	BaseBlog
	Subscribers []*User `gorm:"many2many:user_blogs;"`
}
