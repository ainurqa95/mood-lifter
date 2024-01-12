package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) handleError(chatID int64, err error) {
	log.Println(err)
	messageText := err.Error()
	msg := tgbotapi.NewMessage(chatID, messageText)
	b.client.Send(msg)
}
