package storage

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
	"github.com/jackc/pgx/v5"
)

//go:generate mockery --name=Storage
type Storage interface {
	CreateUser(
		ctx context.Context,
		login string,
		password string,
	) (*models.User, error)

	GetUserByLogin(
		ctx context.Context,
		login string,
	) (*models.User, error)

	CreateOrder(
		ctx context.Context,
		userID int,
		number string,
	) error

	GetOrderByNumber(
		ctx context.Context,
		login string,
	) (*models.Order, error)

	GetOrdersByUserID(
		ctx context.Context,
		userID int,
	) (*[]models.Order, error)

	GetUserByID(
		ctx context.Context,
		userID int,
	) (*models.UserBalance, error)

	GetUserByIDForUpdate(
		ctx context.Context,
		tx pgx.Tx,
		userID int,
	) (*models.UserBalance, error)

	Begin(
		ctx context.Context,
	) (pgx.Tx, error)

	UpdateBalance(
		ctx context.Context,
		tx pgx.Tx,
		sum float32,
		gamblerID int,
	) error

	CreateWithdrawal(
		ctx context.Context,
		tx pgx.Tx,
		userID int,
		sum float32,
		orderID int,
	) error

	GetWithdrawals(
		ctx context.Context,
		userID int,
	) (*[]models.Withdrawal, error)

	UpdateOrder(
		ctx context.Context,
		tx pgx.Tx,
		status string,
		number string,
		accrual *float32,
	) error
}
