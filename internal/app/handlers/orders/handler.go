package orders

import (
	"github.com/KraDM09/gophermart/internal/app/storage"
	"github.com/KraDM09/gophermart/internal/app/validator"
)

type OrderHandler struct {
	store     storage.Storage
	validator validator.Validator
}

func NewHandler(
	store storage.Storage,
	validator validator.Validator,
) *OrderHandler {
	return &OrderHandler{
		store:     store,
		validator: validator,
	}
}
