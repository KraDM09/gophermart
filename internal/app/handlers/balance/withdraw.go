package balance

import (
	"context"
	"encoding/json"
	"github.com/KraDM09/gophermart/internal/constants"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type WithdrawRequest struct {
	Order string  `json:"order" validate:"required,min=9,max=16"`
	Sum   float32 `json:"sum" validate:"required,min=1"`
}

func (h *BalanceHandler) WithdrawHandler(
	ctx context.Context,
	rw http.ResponseWriter,
	r *http.Request,
) {
	var req WithdrawRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(rw, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
		return
	}

	userID := r.Context().Value(constants.ContextUserIDKey).(int)

	tx, err := h.store.Begin(ctx)

	defer tx.Rollback(ctx)

	if err != nil {
		h.logger.Error("Не удалось начать транзакцию")
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	user, err := h.store.GetUserByIDForUpdate(ctx, tx, userID)

	if err != nil {
		h.logger.Error("Не удалось получить пользователя")
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user.Balance-float32(req.Sum) < 0 {
		http.Error(rw, "На счету недостаточно средств", http.StatusPaymentRequired)
		return
	}

	err = h.store.CreateWithdrawal(ctx, tx, userID, req.Sum, req.Order)

	switch {
	case err == pgx.ErrNoRows:
		h.logger.Error("Вывод средств не создан")
		http.Error(rw, "Неверный номер заказа", http.StatusUnprocessableEntity)
		return
	case err != nil:
		h.logger.Error("Не удалось создать вывод средств")
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = h.store.UpdateBalance(ctx, tx, -float32(req.Sum), userID)

	if err != nil {
		h.logger.Error("Не удалось обновить баланс")
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tx.Commit(ctx)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
