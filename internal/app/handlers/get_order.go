package handlers

import (
	"context"
	"encoding/json"
	"github.com/KraDM09/gophermart/internal/app/config"
	"io"
	"net/http"
)

type NumberRequest struct {
	Number string `json:"number" validate:"required,min=6,max=16"`
}

type GetOrderResponse struct {
	Order   string   `json:"order"`
	Status  string   `json:"status"`
	Accrual *float32 `json:"accrual,omitempty"`
}

func (h *UserHandler) GetOrderHandler(
	ctx context.Context,
	rw http.ResponseWriter,
	_ *http.Request,
	number string,
) {
	err := h.validator.Struct(NumberRequest{
		Number: number,
	})

	if err != nil {
		http.Error(rw, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, config.FlagAccrualSystemAddr, nil)

	if err != nil {
		http.Error(rw, "Ошибка при создании нового запроса", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		http.Error(rw, "Ошибка при запросе к сервису расчёта начислений", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(rw, "Ошибка при чтении тела ответа", http.StatusInternalServerError)
		return
	}

	response := &GetOrderResponse{}

	err = json.Unmarshal(buf, response)
	if err != nil {
		http.Error(rw, "Ошибка при декодировании JSON", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(response)
}
