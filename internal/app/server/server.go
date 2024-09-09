package server

import (
	"context"
	"github.com/KraDM09/gophermart/internal/app/validator"
	"net/http"

	"github.com/KraDM09/gophermart/internal/app/access"
	"github.com/KraDM09/gophermart/internal/app/compressor"

	"github.com/KraDM09/gophermart/internal/app/config"
	"github.com/KraDM09/gophermart/internal/app/logger"
	"github.com/KraDM09/gophermart/internal/app/router"
	"github.com/KraDM09/gophermart/internal/app/storage"
)

func Run(
	ctx context.Context,
	store storage.Storage,
	validator validator.Validator,
	r router.Router,
	logger logger.Logger,
	compressor compressor.Compressor,
	access access.Access,
) error {
	if err := logger.Initialize(config.FlagLogLevel); err != nil {
		return err
	}

	validator.Initialize()

	instance := newApp(ctx, store, validator, r, logger, compressor, access)

	instance.logger.Info("Running server", "address", config.FlagRunAddr)
	return http.ListenAndServe(config.FlagRunAddr, instance.webhook(ctx))
}
