package storage

import "context"

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
