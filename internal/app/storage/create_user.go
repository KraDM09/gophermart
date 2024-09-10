package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
)

func (pg PG) CreateUser(
	ctx context.Context,
	login string,
	password string,
) (*models.User, error) {
	var user models.User
	err := pg.pool.QueryRow(ctx,
		`INSERT INTO db_gophermart.users (login, password)
			VALUES ($1, $2) ON CONFLICT (login) DO NOTHING RETURNING id, login, password`,
		login,
		password,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
