package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const commandStart = "start"

func (tg *TelegramBotImpl) handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой комманды")

	switch message.Command() {
	case commandStart:
		msg.Text = "Приветсвую тебя! /start"
		_, err := bot.Send(msg)
		return err
	default:
		_, err := bot.Send(msg)
		return err
	}
}
