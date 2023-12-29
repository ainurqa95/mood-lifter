package telegram

import (
	"context"
	"fmt"
	"github.com/ainurqa95/mood-lifter/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCommand = "start"
)

type Bot struct {
	client *tgbotapi.BotAPI
	cfg    config.Config
}

func NewBot(
	cfg config.Config,
) (*Bot, error) {
	clien, err := tgbotapi.NewBotAPI(cfg.TgCfg.GetToken())
	if err != nil {
		return nil, err
	}
	return &Bot{
		client: clien,
		cfg:    cfg,
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
			err := b.handleCommand(update.Message)
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

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	//authUrl, err := b.InitAuthorization(message.Chat.ID)

	authUrlMessage := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("start messages"))

	_, err := b.client.Send(authUrlMessage)

	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "b.messages.UnknownCommand")
	_, err := b.client.Send(msg)

	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "text")

	_, err := b.client.Send(msg)

	return err
}
