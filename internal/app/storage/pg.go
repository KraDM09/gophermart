package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PG struct {
	pool *pgxpool.Pool
}

func (pg PG) NewStore(ctx context.Context) (*PG, error) {
	pool, err := pgxpool.New(ctx, config.FlagDatabaseDsn)
	if err != nil {
		return nil, err
	}

	return &PG{
		pool: pool,
	}, nil
}

func (pg PG) Begin(
	ctx context.Context,
) (pgx.Tx, error) {
	tx, err := pg.pool.Begin(ctx)

	return tx, err
}
