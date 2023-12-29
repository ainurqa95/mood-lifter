package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	invalidUrlError        = errors.New("Url is invalid")
	unableToSaveError      = errors.New("Unable to save link to Pocket")
	emptyRequestTokenError = errors.New("Empty request token")
	emptyAccessTokenError  = errors.New("Empty accesss token")
)

func (b *Bot) handleError(chatID int64, err error) {
	messageText := err.Error()

	switch err {
	case invalidUrlError:
		messageText = "b.messages.Errors.InvalidURL"
	case unableToSaveError:
		messageText = "b.messages.Errors.UnableToSave"
	case emptyAccessTokenError:
		messageText = "b.messages.EmptyAccessToken"
	case emptyRequestTokenError:
		messageText = "b.messages.EmptyRequestToken"
	default:
		messageText = "b.messages.Errors.Default"
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.client.Send(msg)
}
