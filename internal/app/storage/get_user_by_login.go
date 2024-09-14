package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
	"github.com/jackc/pgx/v5"
)

func (pg PG) GetUserByLogin(
	ctx context.Context,
	login string,
) (*models.User, error) {
	var user models.User
	err := pg.pool.QueryRow(ctx,
		`SELECT id, login, password
				FROM db_gophermart.users
				WHERE login = $1
				LIMIT 1`,
		login,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
