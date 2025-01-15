package repository

//go:generate mockgen -destination ./mock/repository.go -package mock . UserRepository,MessageRepository,ComplimentRepository

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userUUID string, info *model.UserInfo) error
	GetByPeriodWithLimitOffset(ctx context.Context, periodTypes []int, limit int, offset int) ([]model.UserInfo, error)
	UpdatePeriodType(ctx context.Context, periodType int, chatId int64) error
}

type MessageRepository interface {
	Create(ctx context.Context, message model.Message) error
}

type ComplimentRepository interface {
	GetRandom(ctx context.Context) (*model.Compliment, error)
}
