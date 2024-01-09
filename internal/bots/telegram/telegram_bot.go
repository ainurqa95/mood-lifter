package telegram

import (
	"context"
	"fmt"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/ainurqa95/mood-lifter/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCommand = "start"
)

type Bot struct {
	client            *tgbotapi.BotAPI
	cfg               config.Config
	userService       service.UserService
	complimentService service.ComplimentService
}

func NewBot(
	cfg config.Config,
	userService service.UserService,
	complimentService service.ComplimentService,
) (*Bot, error) {
	client, err := tgbotapi.NewBotAPI(cfg.TgCfg.GetToken())
	if err != nil {
		return nil, err
	}
	return &Bot{
		client:            client,
		cfg:               cfg,
		userService:       userService,
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
		fmt.Println(update.Message.Text)
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

		err := b.handleMessage(update.Message)
		fmt.Println("err", err)
		if err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
}

func (b *Bot) handleCommand(ctx context.Context, message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return b.handleStartCommand(ctx, message)
	default:
		return b.handleUnknownCommand(ctx, message)
	}
}

func (b *Bot) handleStartCommand(ctx context.Context, message *tgbotapi.Message) error {
	_, err := b.userService.CreateIfNotExists(ctx, &model.UserInfo{
		UserName: message.Chat.UserName,
		Name:     message.Chat.FirstName,
		ChatID:   message.Chat.ID,
	})
	if err != nil {
		return err
	}
	startMessage := tgbotapi.NewMessage(message.Chat.ID, config.START_MESSAGE)

	_, err = b.client.Send(startMessage)
	if err != nil {
		return err
	}
	compliment, err := b.complimentService.GetRandom(ctx)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf(compliment.Text, message.Chat.FirstName)
	complimentMessage := tgbotapi.NewMessage(message.Chat.ID, msg)
	_, err = b.client.Send(complimentMessage)
	if err != nil {
		return err
	}

	return err
}

func (b *Bot) handleUnknownCommand(ctx context.Context, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "b.messages.UnknownCommand")
	_, err := b.client.Send(msg)

	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "text")

	_, err := b.client.Send(msg)

	return err
}
