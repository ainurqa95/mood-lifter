package service

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
)

type UserService interface {
	CreateIfNotExists(ctx context.Context, info *model.UserInfo) (string, error)
	Get(ctx context.Context, uuid string) (*model.User, error)
}

type ComplimentService interface {
	MassCreate(ctx context.Context) error
	GetRandom(ctx context.Context) (*model.Compliment, error)
}
