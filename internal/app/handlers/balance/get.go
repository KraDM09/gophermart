package balance

import (
	"context"
	"encoding/json"
	"github.com/KraDM09/gophermart/internal/constants"
	"net/http"
)

type UserBalance struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
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
		Current:   user.Balance,
		Withdrawn: user.Withdrawn,
	}

	json.NewEncoder(rw).Encode(balance)
}
