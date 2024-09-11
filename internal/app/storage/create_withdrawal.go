package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (pg PG) CreateWithdrawal(
	ctx context.Context,
	tx pgx.Tx,
	userID int,
	sum float32,
	orderID int,
) error {
	row, err := tx.Exec(ctx,
		`INSERT INTO db_gophermart.withdrawals (user_id, sum, order_id)
				VALUES ($1, $2, $3)
				ON CONFLICT DO NOTHING
				RETURNING id`,
		userID,
		sum,
		orderID,
	)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
