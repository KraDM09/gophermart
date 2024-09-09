package orders

import (
	"context"
	"encoding/json"
	"github.com/KraDM09/gophermart/internal/constants"
	"net/http"
	"time"
)

type Order struct {
	Number     string `json:"number"`
	Status     string `json:"status"`
	UploadedAt string `json:"uploaded_at"`
	Accrual    *int   `json:"accrual"`
}

func (h *OrderHandler) GetHandler(
	ctx context.Context,
	rw http.ResponseWriter,
	r *http.Request,
) {
	userID := r.Context().Value(constants.ContextUserIDKey).(int)

	orders, err := h.store.GetOrdersByUserID(ctx, userID)

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	if len(*orders) == 0 {
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	reps := make([]Order, 0, len(*orders))

	for _, order := range *orders {
		o := Order{
			Number:     order.Number,
			Status:     order.Status,
			UploadedAt: order.UploadedAt.Format(time.RFC3339),
		}

		if order.Accrual != nil {
			accrual := int(*order.Accrual)
			o.Accrual = &accrual
		}

		reps = append(reps, o)
	}

	json.NewEncoder(rw).Encode(reps)
}
