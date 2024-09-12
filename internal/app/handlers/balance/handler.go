package balance

import (
	"github.com/KraDM09/gophermart/internal/app/logger"
	"github.com/KraDM09/gophermart/internal/app/storage"
	"github.com/KraDM09/gophermart/internal/app/validator"
)

type BalanceHandler struct {
	store     storage.Storage
	validator validator.Validator
	logger    logger.Logger
}

func NewHandler(
	store storage.Storage,
	validator validator.Validator,
	logger logger.Logger,
) *BalanceHandler {
	return &BalanceHandler{
		store:     store,
		validator: validator,
		logger:    logger,
	}
}
