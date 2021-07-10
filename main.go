package main

import (
	"github.com/sirupsen/logrus"

	"github.com/squuuze/telegram_gmail_bot/config"
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
}
