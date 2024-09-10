package orders

import (
	"github.com/KraDM09/gophermart/internal/app/models"
	"github.com/KraDM09/gophermart/internal/app/storage"
	"github.com/KraDM09/gophermart/internal/app/validator"
)

type OrderHandler struct {
	store     storage.Storage
	validator validator.Validator
	jobChan   chan models.WorkerPoolJob
}

func NewHandler(
	store storage.Storage,
	validator validator.Validator,
	jobChan chan models.WorkerPoolJob,
) *OrderHandler {
	return &OrderHandler{
		store:     store,
		validator: validator,
		jobChan:   jobChan,
	}
}
