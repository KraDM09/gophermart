package storage

import "context"

func (pg PG) CreateOrder(
	ctx context.Context,
	userID int,
	number string,
) error {
	var id int
	err := pg.pool.QueryRow(ctx,
		`INSERT INTO db_gophermart.orders (user_id, number)
			VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING id`,
		userID,
		number,
	).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}
