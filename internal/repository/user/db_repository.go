package user

import (
	"context"
	"errors"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/ainurqa95/mood-lifter/internal/repository/user/converter"
	repoModel "github.com/ainurqa95/mood-lifter/internal/repository/user/model"
	"github.com/jackc/pgx/v5"
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

func (r *DbUserRepository) GetByPeriodWithLimitOffset(ctx context.Context, periodTypes []int, limit int, offset int) ([]model.UserInfo, error) {
	query := "SELECT id, uid, created_at, chat_id, name, username, period_type FROM users WHERE period_type = ANY($1) ORDER BY id LIMIT $2 OFFSET $3"

	rows, err := r.pool.Query(ctx, query, periodTypes, limit, offset)
	if err != nil {
		return nil, err
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[repoModel.User])
	if errors.Is(err, pgx.ErrNoRows) {
		return []model.UserInfo{}, nil
	}
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return []model.UserInfo{}, nil
	}
	modelInfos := make([]model.UserInfo, len(users))
	for i, user := range users {
		u := converter.ToUserFromRepo(user)
		modelInfos[i] = u.Info
	}

	return modelInfos, nil
}

func (r *DbUserRepository) UpdatePeriodType(ctx context.Context, periodType int, chatId int64) error {
	query := "UPDATE users SET period_type = $1 WHERE chat_id = $2"

	_, err := r.pool.Exec(ctx, query, periodType, chatId)

	return err
}
