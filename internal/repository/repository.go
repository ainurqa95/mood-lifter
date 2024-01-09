package repository

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userUUID string, info *model.UserInfo) error
	Get(ctx context.Context, uuid string) (*model.User, error)
	GetByChatId(ctx context.Context, botId int) (*model.User, error)
}

type ComplimentRepository interface {
	MassCreate(ctx context.Context, compliments []string) error
	GetRandom(ctx context.Context) (*model.Compliment, error)
}
