package user

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DbUserRepository struct {
	pool *pgxpool.Pool
}

func NewDbUserRepository(pool *pgxpool.Pool) *DbUserRepository {
	return &DbUserRepository{
		pool: pool,
	}
}

func (r *DbUserRepository) Create(ctx context.Context, userUUID string, info *model.UserInfo) error {
	query := "INSERT INTO users (uid, chat_id, name, username, created_at) " +
		"VALUES($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING"

	_, err := r.pool.Exec(ctx, query, userUUID, info.ChatID, info.Name, info.UserName, time.Now())

	return err
}

func (r *DbUserRepository) Get(ctx context.Context, uuid string) (*model.User, error) {

	return nil, nil
}

func (r *DbUserRepository) GetByChatId(ctx context.Context, chatId int) (*model.User, error) {
	return nil, nil
}
