package repository

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userUUID string, info *model.UserInfo) error
	GetByLimitOffset(ctx context.Context, limit int, offset int) ([]model.UserInfo, error)
}

type MessageRepository interface {
	Create(ctx context.Context, message model.Message) error
}

type ComplimentRepository interface {
	GetRandom(ctx context.Context) (*model.Compliment, error)
}
