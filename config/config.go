package config

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigT struct {
	BotToken string
	SqliteDb string
	AdminIds []int64
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
