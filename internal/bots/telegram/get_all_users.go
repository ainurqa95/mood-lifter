package telegram

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

// TODO временный метод выпилить
func (b *Bot) handleAllUsersCommand(ctx context.Context, message *tgbotapi.Message) error {
	users, err := b.userService.GetUsersByPeriodWithOffset(ctx, model.AvailableSchedulePeriodTypes, 1000, 0)
	if err != nil {
		return err
	}
	usersStr := make([]string, len(users))
	for i, user := range users {
		usersStr[i] = "@" + user.UserName
	}
	msgText := strings.Join(usersStr, ",")
	startMessage := tgbotapi.NewMessage(message.Chat.ID, msgText)

	_, err = b.client.Send(startMessage)

	return err
}
