package main

import (
	"bytes"
	"encoding/json"
	"errors"
	userHandler "github.com/KraDM09/gophermart/internal/app/handlers"
	balanceHandler "github.com/KraDM09/gophermart/internal/app/handlers/balance"
	ordersHandler "github.com/KraDM09/gophermart/internal/app/handlers/orders"
	ms "github.com/KraDM09/gophermart/internal/app/storage/mocks"
	"github.com/KraDM09/gophermart/internal/app/storage/models"
	mv "github.com/KraDM09/gophermart/internal/app/validator/mocks"
	"github.com/KraDM09/gophermart/internal/constants"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"golang.org/x/net/context"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func Test_get_balance(t *testing.T) {
	t.Run("GetBalance", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		storageProvider.
			On("GetUserByID", mock.Anything, mock.Anything).
			Return(&models.UserBalance{
				ID:        1,
				Balance:   100.0,
				Withdrawn: 50.0,
			}, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/user/balance", nil)
		rr := httptest.NewRecorder()

		h := balanceHandler.NewHandler(storageProvider, validatorProvider, nil)
		ctx := context.WithValue(req.Context(), constants.ContextUserIDKey, 1)
		req = req.WithContext(ctx)

		h.GetHandler(ctx, rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func Test_login(t *testing.T) {
	t.Run("Login", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		storageProvider.
			On("GetUserByLogin", mock.Anything, mock.Anything).
			Return(&models.User{
				ID:       1,
				Login:    "test_user",
				Password: "$2a$10$rxmvdtVOO9mdOdzT1OQbOezRMjkgi83NJuW4LPAq.ZYVT5r9IT/PC",
			}, nil)

		validatorProvider.
			On("Struct", mock.Anything).
			Return(nil)

		loginRequest := userHandler.LoginRequest{
			Login:    "test_user",
			Password: "123456",
		}

		reqBody, err := json.Marshal(loginRequest)
		if err != nil {
			t.Fatalf("Не удалось сериализовать loginRequest в JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()

		h := userHandler.NewHandler(storageProvider, validatorProvider, nil)
		h.LoginHandler(ctx, rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func Test_login_unauthorized(t *testing.T) {
	t.Run("Login", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		storageProvider.
			On("GetUserByLogin", mock.Anything, mock.Anything).
			Return(nil, nil)

		validatorProvider.
			On("Struct", mock.Anything).
			Return(nil)

		loginRequest := userHandler.LoginRequest{
			Login:    "test_user",
			Password: "123456",
		}

		reqBody, err := json.Marshal(loginRequest)
		if err != nil {
			t.Fatalf("Не удалось сериализовать loginRequest в JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()

		h := userHandler.NewHandler(storageProvider, validatorProvider, nil)
		h.LoginHandler(ctx, rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}

func Test_register(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		storageProvider.
			On("GetUserByLogin", mock.Anything, mock.Anything).
			Return(nil, nil)

		storageProvider.
			On("CreateUser", mock.Anything, mock.Anything, mock.Anything).
			Return(&models.User{
				ID:       1,
				Login:    "test_user",
				Password: "$2a$10$rxmvdtVOO9mdOdzT1OQbOezRMjkgi83NJuW4LPAq.ZYVT5r9IT/PC", // bcrypt hash
			}, nil)

		validatorProvider.
			On("Struct", mock.Anything).
			Return(nil)

		registerRequest := userHandler.RegisterRequest{
			Login:    "test_user",
			Password: "123456",
		}

		reqBody, err := json.Marshal(registerRequest)
		if err != nil {
			t.Fatalf("Не удалось сериализовать registerRequest в JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()

		h := userHandler.NewHandler(storageProvider, validatorProvider, nil)
		h.RegisterHandler(ctx, rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func Test_register_conflict(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		storageProvider.
			On("GetUserByLogin", mock.Anything, mock.Anything).
			Return(&models.User{
				ID:       1,
				Login:    "test_user",
				Password: "$2a$10$rxmvdtVOO9mdOdzT1OQbOezRMjkgi83NJuW4LPAq.ZYVT5r9IT/PC", // bcrypt hash
			}, nil)

		validatorProvider.
			On("Struct", mock.Anything).
			Return(nil)

		registerRequest := userHandler.RegisterRequest{
			Login:    "test_user",
			Password: "123456",
		}

		reqBody, err := json.Marshal(registerRequest)
		if err != nil {
			t.Fatalf("Не удалось сериализовать registerRequest в JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()

		h := userHandler.NewHandler(storageProvider, validatorProvider, nil)
		h.RegisterHandler(ctx, rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
	})
}

func Test_register_bad_request(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		validatorProvider.
			On("Struct", mock.Anything).
			Return(errors.New("тут ошибка валидации"))

		registerRequest := userHandler.RegisterRequest{
			Login:    "test_user",
			Password: "123456",
		}

		reqBody, err := json.Marshal(registerRequest)
		if err != nil {
			t.Fatalf("Не удалось сериализовать registerRequest в JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()

		h := userHandler.NewHandler(storageProvider, validatorProvider, nil)
		h.RegisterHandler(ctx, rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func Test_get_orders(t *testing.T) {
	t.Run("GetOrders", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		accural := float32(120.50)

		storageProvider.
			On("GetOrdersByUserID", mock.Anything, mock.Anything).
			Return(&[]models.Order{
				{
					ID:         1,
					Number:     "12345678903",
					Status:     "NEW",
					UploadedAt: time.Now(),
				},
				{
					ID:         2,
					Number:     "62345678903",
					Status:     "PROCESSED",
					UploadedAt: time.Now(),
					Accrual:    &accural,
				},
			}, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/user/orders", nil)
		rr := httptest.NewRecorder()

		h := ordersHandler.NewHandler(storageProvider, validatorProvider, nil)
		ctx := context.WithValue(req.Context(), constants.ContextUserIDKey, 1)
		req = req.WithContext(ctx)
		h.GetHandler(ctx, rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func Test_get_orders_no_content(t *testing.T) {
	t.Run("GetOrders", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		storageProvider.
			On("GetOrdersByUserID", mock.Anything, mock.Anything).
			Return(&[]models.Order{}, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/user/orders", nil)
		rr := httptest.NewRecorder()

		h := ordersHandler.NewHandler(storageProvider, validatorProvider, nil)
		ctx := context.WithValue(req.Context(), constants.ContextUserIDKey, 1)
		req = req.WithContext(ctx)
		h.GetHandler(ctx, rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
	})
}

func Test_withdrawals(t *testing.T) {
	t.Run("GetOrders", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		storageProvider.
			On("GetWithdrawals", mock.Anything, mock.Anything).
			Return(&[]models.Withdrawal{
				{
					Order:       "12345678903",
					Sum:         100.0,
					ProcessedAt: time.Now(),
				},
				{
					Order:       "62345678903",
					Sum:         50.0,
					ProcessedAt: time.Now(),
				},
			}, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/user/withdrawals", nil)
		rr := httptest.NewRecorder()

		h := userHandler.NewHandler(storageProvider, validatorProvider, nil)
		ctx := context.WithValue(req.Context(), constants.ContextUserIDKey, 1)
		req = req.WithContext(ctx)
		h.WithdrawalsHandler(ctx, rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func Test_withdrawals_no_content(t *testing.T) {
	t.Run("GetOrders", func(t *testing.T) {
		storageProvider := new(ms.Storage)
		validatorProvider := new(mv.Validator)

		storageProvider.
			On("GetWithdrawals", mock.Anything, mock.Anything).
			Return(&[]models.Withdrawal{}, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/user/withdrawals", nil)
		rr := httptest.NewRecorder()

		h := userHandler.NewHandler(storageProvider, validatorProvider, nil)
		ctx := context.WithValue(req.Context(), constants.ContextUserIDKey, 1)
		req = req.WithContext(ctx)
		h.WithdrawalsHandler(ctx, rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
	})
}
