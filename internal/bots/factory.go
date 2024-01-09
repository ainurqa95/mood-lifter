package bots

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/bots/telegram"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/service"
)

type BotStarter interface {
	Start(ctx context.Context)
}

func DefineBot(
	cfg config.Config,
	service service.UserService,
	complimentService service.ComplimentService,
) (BotStarter, error) {
	if cfg.BotType == config.TgBot {
		return telegram.NewBot(cfg, service, complimentService)
	}

	return telegram.NewBot(cfg, service, complimentService)
}
