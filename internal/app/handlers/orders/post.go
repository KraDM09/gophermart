package orders

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/models"
	"github.com/KraDM09/gophermart/internal/app/util"
	"github.com/KraDM09/gophermart/internal/constants"
	"io"
	"net/http"
)

type OrderPostRequest struct {
	OrderID string `json:"order_id" validate:"required,min=6,max=16"`
}

func (h *OrderHandler) PostHandler(
	ctx context.Context,
	rw http.ResponseWriter,
	r *http.Request,
) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	number := string(body)
	err = h.validator.Struct(OrderPostRequest{
		OrderID: number,
	})

	if err != nil {
		http.Error(rw, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	if !util.CheckLuna(number) {
		http.Error(rw, "Неверный формат номера заказа", http.StatusUnprocessableEntity)
		return
	}

	order, err := h.store.GetOrderByNumber(ctx, number)

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	userID := r.Context().Value(constants.ContextUserIDKey).(int)

	if order != nil {
		if order.UserID != userID {
			http.Error(rw, "Номер заказа уже был загружен другим пользователем", http.StatusConflict)
			return
		}

		rw.WriteHeader(http.StatusOK)
		return
	}

	err = h.store.CreateOrder(ctx, userID, number)

	h.jobChan <- models.WorkerPoolJob{
		Number: number,
		UserID: userID,
	}

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}
