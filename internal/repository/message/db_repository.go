package message

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DbMessageRepository struct {
	pool *pgxpool.Pool
}

func NewDbMessageRepository(pool *pgxpool.Pool) *DbMessageRepository {
	return &DbMessageRepository{
		pool: pool,
	}
}

func (r *DbMessageRepository) Create(ctx context.Context, message model.Message) error {
	query := "INSERT INTO messages (chat_id, text, created_at) " +
		"VALUES($1, $2, $3)"

	_, err := r.pool.Exec(ctx, query, message.ChatId, message.Text, time.Now())

	return err
}
