package main

import (
	"log"
	"os"

	"./chatwork"
	"./config"
)

func main() {
	token := os.Getenv(config.Config.Token)
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	var message = "テスト"
	var userids = []string{}
	chatwork.CreateTask(message, config.Config.Token, userids, config.Config.Roomid)
	chatwork.Post(message, config.Config.Token, config.Config.Roomid)
}
