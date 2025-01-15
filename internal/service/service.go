package service

//go:generate mockgen -destination ./mock/service.go -package mock . UserService,ComplimentService,MessageService

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
)

type UserService interface {
	CreateIfNotExists(ctx context.Context, info *model.UserInfo) (string, error)
	GetUsersByPeriodWithOffset(ctx context.Context, periodTypes []int, limit int, offset int) ([]model.UserInfo, error)
	UpdatePeriodType(ctx context.Context, periodType int, chatId int64) error
}

type ComplimentService interface {
	GetRandom(ctx context.Context) (*model.Compliment, error)
}

type MessageService interface {
	Create(ctx context.Context, chatId int64, text string) error
}
