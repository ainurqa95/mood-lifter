package telegram

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/mock"
	"github.com/ainurqa95/mood-lifter/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestClientMessagesToBot(t *testing.T) {
	ctrl := gomock.NewController(t)
	userService := mock.NewMockUserService(ctrl)
	complimentService := mock.NewMockComplimentService(ctrl)
	messageService := mock.NewMockMessageService(ctrl)
	client := mock.NewMockBotClient(ctrl)
	bot, err := NewBot(
		client,
		userService,
		complimentService,
		messageService,
	)
	assert.Nil(t, err)
	message := &tgbotapi.Message{
		Chat: &tgbotapi.Chat{
			ID:        123,
			FirstName: "lols",
		},
	}
	ctx := context.TODO()
	userService.EXPECT().CreateIfNotExists(gomock.Any(), gomock.Any()).Return(uuid.NewString(), nil).AnyTimes()
	client.EXPECT().Send(gomock.Any()).Return(tgbotapi.Message{}, nil).AnyTimes()
	messageService.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	complimentService.EXPECT().GetRandom(gomock.Any()).Return(&model.Compliment{}, nil).AnyTimes()
	err = bot.handleStartCommand(ctx, message)
	assert.Nil(t, err)
	err = bot.handleStartCommand(ctx, message)
	assert.Nil(t, err)
}
