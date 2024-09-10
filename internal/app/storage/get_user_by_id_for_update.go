package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
	"github.com/jackc/pgx/v5"
)

func (pg PG) GetUserByIDForUpdate(
	ctx context.Context,
	tx pgx.Tx,
	userID int,
) (*models.UserBalance, error) {
	var user models.UserBalance
	err := tx.QueryRow(ctx,
		`SELECT id, balance
				FROM db_gophermart.users
				WHERE id = $1
				LIMIT 1 FOR UPDATE`,
		userID,
	).Scan(
		&user.ID,
		&user.Balance,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
