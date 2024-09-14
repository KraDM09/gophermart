package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
	"github.com/jackc/pgx/v5"
)

func (pg PG) GetOrderByNumber(
	ctx context.Context,
	number string,
) (*models.Order, error) {
	var order models.Order
	err := pg.pool.QueryRow(ctx,
		`SELECT id, user_id, number, uploaded_at, status, accrual
				FROM db_gophermart.orders
				WHERE number = $1
				LIMIT 1`,
		number,
	).Scan(
		&order.ID,
		&order.UserID,
		&order.Number,
		&order.UploadedAt,
		&order.Status,
		&order.Accrual,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &order, nil
}
