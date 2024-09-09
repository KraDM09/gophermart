package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PG struct {
	pool *pgxpool.Pool
}

func (pg PG) NewStore(ctx context.Context) (*PG, error) {
	pool, err := pgxpool.New(ctx, config.FlagDatabaseDsn)
	if err != nil {
		return nil, err
	}

	return &PG{
		pool: pool,
	}, nil
}

func (pg PG) CreateUser(
	ctx context.Context,
	login string,
	password string,
) (*User, error) {
	var user User
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

func (pg PG) Bootstrap(
	ctx context.Context,
) error {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.schemata
			WHERE schema_name = 'db_gophermart'
		)
	`

	err := pg.pool.QueryRow(ctx, query).Scan(&exists)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	tx, err := pg.pool.Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
        CREATE schema IF NOT EXISTS db_gophermart
    `)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS db_gophermart.users
		(
			id          INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			create_dttm TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			login       TEXT UNIQUE,
			password    TEXT,
            balance     DECIMAL(10, 1) DEFAULT 0
		);

		CREATE TYPE db_gophermart.order_statuses AS ENUM (
			'NEW',
			'PROCESSING',
			'INVALID',
			'PROCESSED');
		
		CREATE TABLE IF NOT EXISTS db_gophermart.orders
		(
			id          INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			user_id     INT REFERENCES db_gophermart.users (id) NOT NULL,
			number      TEXT                                    NOT NULL UNIQUE,
			uploaded_at TIMESTAMP WITH TIME ZONE     DEFAULT NOW(),
			status      db_gophermart.order_statuses DEFAULT 'NEW'::db_gophermart.order_statuses,
			accrual     NUMERIC(10, 1)               DEFAULT NULL
		);

	    CREATE TABLE IF NOT EXISTS db_gophermart.withdrawals
		(
			id           INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			order_id     INT REFERENCES db_gophermart.orders (id) NOT NULL,
			user_id      INT REFERENCES db_gophermart.users (id)  NOT NULL,
			sum          INTEGER                                  NOT NULL,
			processed_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);
    `)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (pg PG) GetUserByLogin(
	ctx context.Context,
	login string,
) (*User, error) {
	var user User
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

func (pg PG) GetOrderByNumber(
	ctx context.Context,
	number string,
) (*Order, error) {
	var order Order
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

func (pg PG) GetOrdersByUserID(
	ctx context.Context,
	userID int,
) (*[]Order, error) {
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

	orders := make([]Order, 0)

	for rows.Next() {
		var order Order
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

func (pg PG) GetUserByID(
	ctx context.Context,
	userID int,
) (*UserBalance, error) {
	var user UserBalance
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

func (pg PG) GetUserByIDForUpdate(
	ctx context.Context,
	tx pgx.Tx,
	userID int,
) (*UserBalance, error) {
	var user UserBalance
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

func (pg PG) Begin(
	ctx context.Context,
) (pgx.Tx, error) {
	tx, err := pg.pool.Begin(ctx)

	return tx, err
}

func (pg PG) UpdateBalance(
	ctx context.Context,
	tx pgx.Tx,
	sum int,
	gamblerID int,
) error {
	query := `UPDATE db_gophermart.users
				SET balance = users.balance - $1
				WHERE id = $2`

	_, err := tx.Exec(
		ctx,
		query,
		sum,
		gamblerID,
	)

	return err
}

func (pg PG) CreateWithdrawal(
	ctx context.Context,
	tx pgx.Tx,
	userID int,
	sum int,
	orderID int,
) error {
	row, err := tx.Exec(ctx,
		`INSERT INTO db_gophermart.withdrawals (user_id, sum, order_id)
				VALUES ($1, $2, $3)
				ON CONFLICT DO NOTHING
				RETURNING id`,
		userID,
		sum,
		orderID,
	)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (pg PG) GetWithdrawals(
	ctx context.Context,
	userID int,
) (*[]Withdrawal, error) {
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

	withdrawals := make([]Withdrawal, 0)

	for rows.Next() {
		var withdrawal Withdrawal
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
