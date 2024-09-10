package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
)

func (pg PG) GetOrdersByUserID(
	ctx context.Context,
	userID int,
) (*[]models.Order, error) {
	rows, err := pg.pool.Query(ctx,
		`SELECT number, uploaded_at, status, accrual
				FROM db_gophermart.orders
				WHERE user_id = $1
				ORDER BY uploaded_at DESC`,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := make([]models.Order, 0)

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.Number,
			&order.UploadedAt,
			&order.Status,
			&order.Accrual,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &orders, nil
}
