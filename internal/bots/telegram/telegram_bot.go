package telegram

import (
	"context"
	"fmt"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/ainurqa95/mood-lifter/internal/service"
	"github.com/ainurqa95/mood-lifter/internal/service/compliment"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const (
	startCommand       = "start"
	getAllUsersCommand = "getAllUsers"
)

type BotClient interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
}

type Bot struct {
	client            BotClient
	cfg               config.Config
	userService       service.UserService
	complimentService service.ComplimentService
	messageService    service.MessageService
}

func NewBot(
	client BotClient,
	userService service.UserService,
	complimentService service.ComplimentService,
	messageService service.MessageService,
) (*Bot, error) {
	return &Bot{
		client:            client,
		userService:       userService,
		messageService:    messageService,
		complimentService: complimentService,
	}, nil
}

func (b *Bot) Start(ctx context.Context) {
	updates := b.initUpdates()

	b.handleUpdates(ctx, updates)
}

func (b *Bot) initUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.client.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		select {
		case <-ctx.Done():
			break
		default:
		}
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := b.handleCommand(ctx, update.Message)
			if err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		err := b.handleMessage(ctx, update.Message)
		if err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
}

func (b *Bot) handleCommand(ctx context.Context, message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return b.handleStartCommand(ctx, message)
	case getAllUsersCommand:
		return b.handleAllUsersCommand(ctx, message)
	default:
		return b.handleUnknownCommand(ctx, message)
	}
}

func (b *Bot) handleStartCommand(ctx context.Context, message *tgbotapi.Message) error {
	err := b.createUser(ctx, message)
	if err != nil {
		return err
	}
	startMessage := tgbotapi.NewMessage(message.Chat.ID, config.START_MESSAGE)

	_, err = b.client.Send(startMessage)
	if err != nil {
		return err
	}
	return b.SendCompliment(ctx, message.Chat.FirstName, message.Chat.ID)
}

func (b *Bot) handleAllUsersCommand(ctx context.Context, message *tgbotapi.Message) error {
	users, err := b.userService.GetUsersByOffset(ctx, compliment.DEFAULT_LIMIT, 0)
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

func (b *Bot) handleMessage(ctx context.Context, message *tgbotapi.Message) error {
	err := b.createUser(ctx, message)
	if err != nil {
		return err
	}

	return b.SendCompliment(ctx, message.Chat.FirstName, message.Chat.ID)
}

func (b *Bot) SendCompliment(ctx context.Context, name string, chatId int64) error {
	randomCompliment, err := b.complimentService.GetRandom(ctx)
	if err != nil {
		return err
	}
	text := fmt.Sprintf(randomCompliment.Text, name)
	complimentMessage := tgbotapi.NewMessage(chatId, text)
	_, err = b.client.Send(complimentMessage)
	if err != nil {
		return fmt.Errorf("ошибки при отправке комплимента %s %v", name, err)
	}

	return b.messageService.Create(ctx, chatId, text)
}

func (b *Bot) createUser(ctx context.Context, message *tgbotapi.Message) error {
	_, err := b.userService.CreateIfNotExists(ctx, &model.UserInfo{
		UserName: message.Chat.UserName,
		Name:     message.Chat.FirstName,
		ChatID:   message.Chat.ID,
	})
	return err
}

func (b *Bot) handleUnknownCommand(ctx context.Context, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "UnknownCommand")
	_, err := b.client.Send(msg)

	return err
}
