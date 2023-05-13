package config

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigT struct {
	BotToken           string
	SqliteDb           string
	AllowedIds         []int64
	AdminIds           []int64
	UpdateIntervalHrs  int
	AllowedErrorMargin int
	TGMaxMessageLength int
}

var Config *ConfigT

func Init() {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Config)
	if err != nil {
		log.Fatal("[!] Can't read config.json", err)
	}
}
