package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (pg PG) UpdateOrder(
	ctx context.Context,
	tx pgx.Tx,
	status string,
	number string,
	accrual *float32,
) error {
	query := `UPDATE db_gophermart.orders
				SET status = $1, accrual = $3
				WHERE number = $2`

	_, err := tx.Exec(
		ctx,
		query,
		status,
		number,
		accrual,
	)

	return err
}
