package service

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
)

type UserService interface {
	CreateIfNotExists(ctx context.Context, info *model.UserInfo) (string, error)
	GetUsersByOffset(ctx context.Context, limit int, offset int) ([]model.UserInfo, error)
}

type ComplimentService interface {
	GetRandom(ctx context.Context) (*model.Compliment, error)
}

type MessageService interface {
	Create(ctx context.Context, chatId int64, text string) error
}
