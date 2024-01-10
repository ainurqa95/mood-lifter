package bots

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/bots/telegram"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/service"
)

type BotManager interface {
	Start(ctx context.Context)
	SendCompliment(ctx context.Context, name string, chatId int64) error
}

func DefineBot(
	cfg config.Config,
	service service.UserService,
	complimentService service.ComplimentService,
	messageService service.MessageService,
) (BotManager, error) {
	if cfg.BotType == config.TgBot {
		return telegram.NewBot(cfg, service, complimentService, messageService)
	}

	return telegram.NewBot(cfg, service, complimentService, messageService)
}
