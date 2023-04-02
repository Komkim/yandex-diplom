package main

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"os"
	"os/signal"
	"syscall"
	"yandex-diplom/config"
	accrualserver "yandex-diplom/internal/accrual/server"
	accrualrouter "yandex-diplom/internal/accrual/server/http"
	"yandex-diplom/internal/accrual/service"
	mylogger "yandex-diplom/pkg/logger"
)

func main() {
	logger := mylogger.NewLogger()
	logger.Debug().Str("accrual", "start")

	ctx, cencel := context.WithCancel(context.Background())

	cfg, err := config.NewAccrualConfig()
	if err != nil {
		logger.Error().Err(err)
		return
	}

	err = startMigration(cfg.DatabaseDSN)
	if err != nil {
		logger.Error().Err(err)
	}

	db, err := newDb(ctx, cfg.DatabaseDSN)
	if err != nil {
		logger.Error().Err(err)
		return
	}

	srv := service.NewSerivces(logger.Log(), db)
	r := accrualrouter.NewAccrualRouter(cfg, srv, logger.Log())
	s := accrualserver.NewServer(cfg, logger.Log(), r.Init())

	go s.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := s.GetServer().Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("Server forced to shutdown:")
	}

	defer cencel()
}

func newDb(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
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
