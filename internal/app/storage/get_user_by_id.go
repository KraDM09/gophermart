package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
	"github.com/jackc/pgx/v5"
)

func (pg PG) GetUserByID(
	ctx context.Context,
	userID int,
) (*models.UserBalance, error) {
	var user models.UserBalance
	err := pg.pool.QueryRow(ctx,
		`SELECT id,
				   balance,
				   (SELECT COALESCE(SUM(sum), 0)
					FROM db_gophermart.withdrawals
					WHERE id = $1) AS withdrawn
			FROM db_gophermart.users
			WHERE id = $1
			LIMIT 1;`,
		userID,
	).Scan(
		&user.ID,
		&user.Balance,
		&user.Withdrawn,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
