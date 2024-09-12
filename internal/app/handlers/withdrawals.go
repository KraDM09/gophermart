package handlers

import (
	"context"
	"encoding/json"
	"github.com/KraDM09/gophermart/internal/constants"
	"net/http"
	"time"
)

type Withdrawal struct {
	Order        string  `json:"order"`
	Sum          float32 `json:"sum"`
	Processed_at string  `json:"processed_at"`
}

func (h *UserHandler) WithdrawalsHandler(
	ctx context.Context,
	rw http.ResponseWriter,
	r *http.Request,
) {
	userID := r.Context().Value(constants.ContextUserIDKey).(int)
	withdrawals, err := h.store.GetWithdrawals(ctx, userID)

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	if len(*withdrawals) == 0 {
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	response := make([]Withdrawal, 0, len(*withdrawals))

	for _, withdrawal := range *withdrawals {
		response = append(response, Withdrawal{
			Order:        withdrawal.Order,
			Sum:          withdrawal.Sum,
			Processed_at: withdrawal.Processed_at.Format(time.RFC3339),
		})
	}

	json.NewEncoder(rw).Encode(response)
}
