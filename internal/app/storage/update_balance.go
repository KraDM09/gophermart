package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (pg PG) UpdateBalance(
	ctx context.Context,
	tx pgx.Tx,
	sum float32,
	gamblerID int,
) error {
	query := `UPDATE db_gophermart.users
				SET balance = users.balance + $1
				WHERE id = $2`

	_, err := tx.Exec(
		ctx,
		query,
		sum,
		gamblerID,
	)

	return err
}
