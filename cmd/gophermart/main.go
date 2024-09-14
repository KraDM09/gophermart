package main

import (
	"context"
	"fmt"
	"github.com/KraDM09/gophermart/internal/app/access"
	"github.com/KraDM09/gophermart/internal/app/compressor"
	"github.com/KraDM09/gophermart/internal/app/config"
	"github.com/KraDM09/gophermart/internal/app/logger"
	"github.com/KraDM09/gophermart/internal/app/router"
	"github.com/KraDM09/gophermart/internal/app/server"
	"github.com/KraDM09/gophermart/internal/app/storage"
	"github.com/KraDM09/gophermart/internal/app/validator"
)

func getStorage(
	ctx context.Context,
) (storage.Storage, error) {
	pg, err := storage.PG{}.NewStore(ctx)
	if err != nil {
		return nil, err
	}

	err = pg.Bootstrap(ctx)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

func main() {
	config.ParseFlags()

	ctx := context.Background()
	store, err := getStorage(ctx)
	if err != nil {
		panic(fmt.Errorf("не удалось получить доступ к хранилищу %w", err))
	}

	r := &router.ChiRouter{}
	log := &logger.ZapLogger{}
	c := &compressor.GzipCompressor{}
	a := &access.Cookie{}
	v := &validator.V10Validator{}

	if err := server.Run(ctx, store, v, r, log, c, a); err != nil {
		panic(fmt.Errorf("ошибка во время старта сервиса %w", err))
	}
}
