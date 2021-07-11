package main

import (
	"github.com/sirupsen/logrus"

	"github.com/squuuze/telegram_gmail_bot/config"
	"github.com/squuuze/telegram_gmail_bot/telegram"
)

func main() {
	cfg := config.Get()
	log := logrus.New()
	logFormatter := logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	}

	log.SetFormatter(&logFormatter)
	log.SetLevel(logrus.DebugLevel)
	log.Printf("Provided config: %v", cfg.String())

	newTgBot := telegram.NewTelegramBot(cfg.TelegramApiToken, log)
	if err := newTgBot.Start(); err != nil {
		log.Panic(err)
	}
}
