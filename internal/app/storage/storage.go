package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserBalance struct {
	ID        int     `json:"id"`
	Balance   float64 `json:"balance"`
	Withdrawn int     `json:"withdrawn"`
}

type Order struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Number     string    `json:"number"`
	UploadedAt time.Time `json:"uploaded_at"`
	Status     string    `json:"status"`
	Accrual    *float64  `json:"accrual"`
}

type Withdrawal struct {
	Number       string    `json:"number"`
	Sum          int       `json:"sum"`
	Processed_at time.Time `json:"processed_at"`
}

//go:generate mockery --name=Storage
type Storage interface {
	CreateUser(
		ctx context.Context,
		login string,
		password string,
	) (*User, error)

	GetUserByLogin(
		ctx context.Context,
		login string,
	) (*User, error)

	CreateOrder(
		ctx context.Context,
		userID int,
		number string,
	) error

	GetOrderByNumber(
		ctx context.Context,
		login string,
	) (*Order, error)

	GetOrdersByUserID(
		ctx context.Context,
		userID int,
	) (*[]Order, error)

	GetUserByID(
		ctx context.Context,
		userID int,
	) (*UserBalance, error)

	GetUserByIDForUpdate(
		ctx context.Context,
		tx pgx.Tx,
		userID int,
	) (*UserBalance, error)

	Begin(
		ctx context.Context,
	) (pgx.Tx, error)

	UpdateBalance(
		ctx context.Context,
		tx pgx.Tx,
		sum int,
		gamblerID int,
	) error

	CreateWithdrawal(
		ctx context.Context,
		tx pgx.Tx,
		userID int,
		sum int,
		orderID int,
	) error

	GetWithdrawals(
		ctx context.Context,
		userID int,
	) (*[]Withdrawal, error)
}
