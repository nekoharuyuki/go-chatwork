package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Roomid string
	Token  string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Roomid: cfg.Section("chatwork").Key("roomid").String(),
		Token:  cfg.Section("chatwork").Key("token").String(),
	}
}
