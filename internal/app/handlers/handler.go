package handlers

import (
	"github.com/KraDM09/gophermart/internal/app/storage"
	"github.com/KraDM09/gophermart/internal/app/validator"
)

type UserHandler struct {
	store     storage.Storage
	validator validator.Validator
}

func NewHandler(
	store storage.Storage,
	validator validator.Validator,
) *UserHandler {
	return &UserHandler{
		store:     store,
		validator: validator,
	}
}
