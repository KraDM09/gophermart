package balance

import (
	"github.com/KraDM09/gophermart/internal/app/storage"
	"github.com/KraDM09/gophermart/internal/app/validator"
)

type BalanceHandler struct {
	store     storage.Storage
	validator validator.Validator
}

func NewHandler(
	store storage.Storage,
	validator validator.Validator,
) *BalanceHandler {
	return &BalanceHandler{
		store:     store,
		validator: validator,
	}
}
