package main

import (
	"context"
	"database/sql"
	"os"
	"yandex-diplom/internal/infrastructure/server/auth"
	"yandex-diplom/internal/updateorder"
	mylogger "yandex-diplom/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"os/signal"
	"syscall"
	"yandex-diplom/config"
	"yandex-diplom/internal/application"
	"yandex-diplom/internal/infrastructure/server"
	router "yandex-diplom/internal/infrastructure/server/http"

	//_ "github.com/jackc/pgx/v5"
	_ "yandex-diplom/storage/postgres/migrations/gophermart"

	"github.com/pressly/goose/v3"
)

func main() {
	logger := mylogger.NewLogger()
	logger.Debug().Str("server", "start")

	ctx, cencel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error().Err(err)
	}

	logger.Info().Msgf("Config %v", cfg)

	err = startMigration(cfg.Server.DatabaseDSN)
	if err != nil {
		logger.Error().Err(err)
	}

	db, err := newDB(ctx, cfg.Server.DatabaseDSN)
	if err != nil {
		log.Err(err)
		return
	}

	a := auth.NewMemoryAuth()
	service := application.NewServices(logger.Log(), db)
	r := router.NewRouter(&cfg.Server, service, a, logger.Log())
	srv := server.NewServer(&cfg.HTTP, logger.Log(), r.Init())
	go srv.Start()

	accrual := updateorder.NewAccrual(&cfg.Server, logger.Log(), service)
	go accrual.Start(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.GetServer().Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("Server forced to shutdown:")
	}

	defer cencel()
}

func startMigration(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	err = goose.Up(db, "/var")
	if err != nil {
		return err
	}
	return nil
}

func newDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
