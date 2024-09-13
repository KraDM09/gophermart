package workerpool

import (
	"context"
	"encoding/json"
	"github.com/KraDM09/gophermart/internal/app/logger"
	"github.com/KraDM09/gophermart/internal/app/models"
	"github.com/KraDM09/gophermart/internal/app/storage"
	"github.com/KraDM09/gophermart/internal/constants"
	"io"
	"net/http"
	"strconv"
	"time"
)

type WorkerPool struct {
	store  storage.Storage
	logger logger.Logger
}

func NewHandler(
	store storage.Storage,
	logger logger.Logger,
) *WorkerPool {
	return &WorkerPool{
		store:  store,
		logger: logger,
	}
}

func (wp WorkerPool) Worker(
	ctx context.Context,
	id int,
	jobs chan models.WorkerPoolJob,
) {

	for job := range jobs {
		wp.logger.Info("рабочий id запущен", "id", strconv.Itoa(id))

		err := wp.updateInfo(ctx, &job)

		if err == nil {
			wp.logger.Info("заказ обработан", "order", job.Number)
			continue
		}

		if err.Code == http.StatusTooManyRequests {
			time.Sleep(60 * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}

		wp.logger.Error(err.Message)
		jobs <- job

		continue
	}
}

type CustomError struct {
	Code    int
	Message string
}

func (wp WorkerPool) updateInfo(
	ctx context.Context,
	job *models.WorkerPoolJob,
) *CustomError {
	response, err := GetOrderHandler(context.Background(), job.Number)

	if err != nil {
		return &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "ошибка при получении заказа" + err.Error(),
		}
	}

	if response.StatusCode == 429 {
		return &CustomError{
			Code:    http.StatusTooManyRequests,
			Message: "превышено количество запросов к сервису",
		}
	}

	if response.StatusCode != 200 {
		return &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "некорректный ответ от API status code: " + strconv.Itoa(response.StatusCode),
		}
	}

	if response.Status == constants.LoyaltyOrderStatusRegistered {
		return &CustomError{
			Code:    http.StatusNotAcceptable,
			Message: "заказ не обработан",
		}
	}

	tx, err := wp.store.Begin(ctx)

	if err != nil {
		return &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "ошибка начала транзакции " + err.Error(),
		}
	}

	defer tx.Rollback(ctx)

	if response.Accrual != nil {
		err = wp.store.UpdateBalance(ctx, tx, *response.Accrual, job.UserID)

		if err != nil {
			return &CustomError{
				Code:    http.StatusInternalServerError,
				Message: "ошибка обновления баланса" + err.Error(),
			}
		}
	}

	err = wp.store.UpdateOrder(ctx, tx, response.Status, job.Number, response.Accrual)

	if err != nil {
		return &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "ошибка обновления заказа" + err.Error(),
		}
	}

	err = tx.Commit(ctx)

	if err != nil {
		return &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "ошибка коммита транзакции" + err.Error(),
		}
	}

	return nil
}

type GetOrderResponse struct {
	StatusCode int      `json:"-"`
	Order      string   `json:"order"`
	Status     string   `json:"status"`
	Accrual    *float32 `json:"accrual,omitempty"`
}

func GetOrderHandler(
	ctx context.Context,
	number string,
) (*GetOrderResponse, error) {
	//url := fmt.Sprintf("%s/api/orders/%s", config.FlagAccrualSystemAddr, number)
	url := "https://7b525102-c05c-4a37-9073-76cc1430afd1.mock.pstmn.io"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &GetOrderResponse{
		StatusCode: resp.StatusCode,
	}

	err = json.Unmarshal(buf, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
