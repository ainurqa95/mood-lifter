package bots

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/bots/telegram"
	"github.com/ainurqa95/mood-lifter/internal/config"
)

type BotStarter interface {
	Start(ctx context.Context)
}

func DefineBot(cfg config.Config) (BotStarter, error) {
	if cfg.BotType == config.TgBot {
		return telegram.NewBot(cfg)
	}

	return telegram.NewBot(cfg)
}
