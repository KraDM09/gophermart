package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
)

func (pg PG) GetWithdrawals(
	ctx context.Context,
	userID int,
) (*[]models.Withdrawal, error) {
	rows, err := pg.pool.Query(ctx,
		`SELECT o.number, w.sum, w.processed_at
				FROM db_gophermart.withdrawals w
						 JOIN db_gophermart.orders o on o.id = w.order_id
				WHERE w.user_id = $1;`,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	withdrawals := make([]models.Withdrawal, 0)

	for rows.Next() {
		var withdrawal models.Withdrawal
		err = rows.Scan(
			&withdrawal.Number,
			&withdrawal.Sum,
			&withdrawal.Processed_at,
		)
		if err != nil {
			return nil, err
		}
		withdrawals = append(withdrawals, withdrawal)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &withdrawals, nil
}
