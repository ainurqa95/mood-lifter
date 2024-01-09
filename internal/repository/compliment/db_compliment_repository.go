package compliment

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DbComplimentRepository struct {
	pool *pgxpool.Pool
}

func NewDbComplimentRepository(pool *pgxpool.Pool) *DbComplimentRepository {
	return &DbComplimentRepository{
		pool: pool,
	}
}

func (r *DbComplimentRepository) MassCreate(ctx context.Context, compliments []string) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	query := "CREATE TEMPORARY TABLE IF NOT EXISTS _tmp_compliments (LIKE compliments INCLUDING ALL) ON COMMIT DELETE ROWS"
	_, err = tx.Exec(ctx, query)
	if err != nil {
		return err
	}
	txsIterate := func(i int) ([]interface{}, error) {
		return []interface{}{
			compliments[i],
			time.Now(),
		}, nil
	}
	_, err = tx.CopyFrom(ctx, pgx.Identifier{"_tmp_compliments"},
		[]string{
			"text",
			"created_at",
		},
		pgx.CopyFromSlice(len(compliments), txsIterate),
	)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "INSERT INTO compliments SELECT * from _tmp_compliments ON CONFLICT DO NOTHING")
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *DbComplimentRepository) GetRandom(ctx context.Context) (*model.Compliment, error) {
	query := "SELECT text, created_at FROM compliments ORDER BY random() LIMIT 1"

	row := r.pool.QueryRow(ctx, query)
	var compliment model.Compliment
	err := row.Scan(&compliment.Text, &compliment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &compliment, nil
}
