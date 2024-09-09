package balance

import (
	"context"
	"encoding/json"
	"github.com/KraDM09/gophermart/internal/constants"
	"net/http"
)

type UserBalance struct {
	Balance   float64 `json:"balance"`
	Withdrawn int     `json:"withdrawn"`
}

func (h *BalanceHandler) GetHandler(
	ctx context.Context,
	rw http.ResponseWriter,
	r *http.Request,
) {
	userID := r.Context().Value(constants.ContextUserIDKey).(int)
	user, err := h.store.GetUserByID(ctx, userID)

	if err != nil {
		http.Error(rw, "Не удалось получить пользователя", http.StatusInternalServerError)
		return
	}

	balance := UserBalance{
		Balance:   user.Balance,
		Withdrawn: user.Withdrawn,
	}

	json.NewEncoder(rw).Encode(balance)
}
