package db

import (
	"blogs_subscriber/config"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TgId      int64 `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Username  string
}

var DB *gorm.DB

func Init() {
	log.Printf("[*] Initializing %s.", config.Config.SqliteDb)
	db, err := gorm.Open(sqlite.Open(config.Config.SqliteDb), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	db.AutoMigrate(&User{})
}
