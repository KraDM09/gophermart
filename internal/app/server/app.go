package server

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/access"
	"github.com/KraDM09/gophermart/internal/app/compressor"
	"github.com/KraDM09/gophermart/internal/app/handlers"
	"github.com/KraDM09/gophermart/internal/app/handlers/balance"
	"github.com/KraDM09/gophermart/internal/app/handlers/orders"
	"github.com/KraDM09/gophermart/internal/app/logger"
	"github.com/KraDM09/gophermart/internal/app/models"
	"github.com/KraDM09/gophermart/internal/app/router"
	"github.com/KraDM09/gophermart/internal/app/storage"
	"github.com/KraDM09/gophermart/internal/app/validator"
	"github.com/KraDM09/gophermart/internal/app/workerpool"
	"github.com/go-chi/chi"
	"net/http"
)

// app инкапсулирует в себя все зависимости и логику приложения
type app struct {
	store      storage.Storage
	validator  validator.Validator
	router     router.Router
	logger     logger.Logger
	compressor compressor.Compressor
	access     access.Access
	jobsChan   chan models.WorkerPoolJob
}

// newApp принимает на вход внешние зависимости приложения и возвращает новый объект app
func newApp(
	ctx context.Context,
	store storage.Storage,
	validator validator.Validator,
	router router.Router,
	logger logger.Logger,
	compressor compressor.Compressor,
	access access.Access,
) *app {
	instance := &app{
		store:      store,
		validator:  validator,
		router:     router,
		logger:     logger,
		compressor: compressor,
		access:     access,
		jobsChan:   make(chan models.WorkerPoolJob, 5),
	}

	go instance.startWorkers(ctx, 3)

	return instance
}

func (a *app) webhook(
	ctx context.Context,
) router.Router {
	a.router.Use(a.logger.RequestLogger)
	a.router.Use(a.compressor.RequestCompressor)

	userHandler := handlers.NewHandler(a.store, a.validator)
	ordersHandler := orders.NewHandler(a.store, a.validator, a.jobsChan)
	balanceHandler := balance.NewHandler(a.store, a.validator, a.logger)

	a.router.Post("/api/user/login", func(rw http.ResponseWriter, r *http.Request) {
		userHandler.LoginHandler(ctx, rw, r)
	})

	a.router.Post("/api/user/register", func(rw http.ResponseWriter, r *http.Request) {
		userHandler.RegisterHandler(ctx, rw, r)
	})

	a.router.Group(func(r chi.Router) {
		r.Use(a.access.Control)

		r.Get("/api/user/orders", func(rw http.ResponseWriter, r *http.Request) {
			ordersHandler.GetHandler(ctx, rw, r)
		})

		r.Post("/api/user/orders", func(rw http.ResponseWriter, r *http.Request) {
			ordersHandler.PostHandler(ctx, rw, r)
		})

		r.Get("/api/user/balance", func(rw http.ResponseWriter, r *http.Request) {
			balanceHandler.GetHandler(ctx, rw, r)
		})

		r.Post("/api/user/balance/withdraw", func(rw http.ResponseWriter, r *http.Request) {
			balanceHandler.WithdrawHandler(ctx, rw, r)
		})

		r.Get("/api/user/withdrawals", func(rw http.ResponseWriter, r *http.Request) {
			userHandler.WithdrawalsHandler(ctx, rw, r)
		})
	})

	return a.router
}

func (a *app) startWorkers(
	ctx context.Context,
	quantity int,
) {
	wp := workerpool.NewHandler(a.store, a.logger)

	for w := 1; quantity <= 3; w++ {
		go wp.Worker(ctx, w, a.jobsChan)
	}

	close(a.jobsChan)
}
