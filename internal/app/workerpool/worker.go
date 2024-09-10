package workerpool

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/KraDM09/gophermart/internal/app/config"
<<<<<<< Updated upstream
=======
	"github.com/KraDM09/gophermart/internal/app/logger"
>>>>>>> Stashed changes
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

		response, err := GetOrderHandler(context.Background(), job.Number)

		if err != nil {
			wp.logger.Error(err.Error())
			time.Sleep(1 * time.Second)
			jobs <- job

			continue
		}

		if response.StatusCode == 429 {
			wp.logger.Error("превышено количество запросов к сервису")
			time.Sleep(60 * time.Second)
			jobs <- job

			continue
		}

		if response.StatusCode != 200 {
			wp.logger.Error("некорректный ответ от API")
			time.Sleep(1 * time.Second)
			jobs <- job

			continue
		}

		if response.Status == constants.LOYALTY_ORDER_STATUS_REGISTERED {
			wp.logger.Error("заказ  to не обработан")
			time.Sleep(1 * time.Second)
			jobs <- job

			continue
		}

		tx, err := wp.store.Begin(ctx)

		if err != nil {
			wp.logger.Error("ошибка начала транзакции")
			time.Sleep(1 * time.Second)
			jobs <- job

			continue
		}

		defer tx.Rollback(ctx)

		if response.Accrual != nil {
			err = wp.store.UpdateBalance(ctx, tx, *response.Accrual, job.UserID)

			if err != nil {
				wp.logger.Error("ошибка обновления баланса")
				time.Sleep(1 * time.Second)
				jobs <- job

				continue
			}
		}

		err = wp.store.UpdateOrder(ctx, tx, response.Status, job.Number, response.Accrual)

		if err != nil {
			wp.logger.Error("ошибка обновления заказа")
			time.Sleep(1 * time.Second)
			jobs <- job

			continue
		}

		err = tx.Commit(ctx)
	}
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
	url := fmt.Sprintf("%s/api/orders/%s", config.FlagAccrualSystemAddr, number)
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
