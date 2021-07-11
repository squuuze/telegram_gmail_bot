package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/sirupsen/logrus"
)

type TelegramBot interface {
	Start() error
	GetMessageChan() chan []byte
	Stop() error
}

type TelegramBotImpl struct {
	tgToken  string
	msgChan  chan []byte
	stopChan chan struct{}
	logger   *logrus.Logger
}

func (tb *TelegramBotImpl) Start() error {
	bot, err := tgbotapi.NewBotAPI(tb.tgToken)
	if err != nil {
		return err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	bot.Debug = true
	tb.logger.Printf("Authorized on account %s", bot.Self.UserName)

	tb.handleUpdates(bot, updates)
	return nil
}

func (tb *TelegramBotImpl) handleUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			tb.handleCommand(bot, update.Message)
			continue
		}

		tb.logger.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func (tb *TelegramBotImpl) GetMessageChan() chan []byte {
	return tb.msgChan
}

func (tb *TelegramBotImpl) Stop() error {
	tb.stopChan <- struct{}{}
	return nil
}

func NewTelegramBot(tgToken string, logger *logrus.Logger) TelegramBot {
	return &TelegramBotImpl{
		tgToken:  tgToken,
		msgChan:  make(chan []byte),
		stopChan: make(chan struct{}),
		logger:   logger,
	}
}
