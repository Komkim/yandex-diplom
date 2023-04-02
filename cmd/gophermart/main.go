package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"yandex-diplom/internal/infrastructure/server/auth"
	"yandex-diplom/pkg/logger"

	"os/signal"
	"syscall"
	"yandex-diplom/config"
	"yandex-diplom/internal/application"
	"yandex-diplom/internal/infrastructure/server"
	router "yandex-diplom/internal/infrastructure/server/http"

	//_ "github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
	_ "yandex-diplom/storage/postgres/migrations/gophermart"
)

func main() {
	logger := mylogger.NewLogger()
	logger.Debug().Str("server", "start")

	ctx, cencel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error().Err(err)
	}

	err = startMigration(cfg.Server.DatabaseDSN)
	if err != nil {
		logger.Error().Err(err)
	}

	a := auth.NewMemoryAuth()
	service := application.NewServices(ctx, &cfg.Server, logger.Log())
	r := router.NewRouter(&cfg.Server, service, a, logger.Log())
	srv := server.NewServer(&cfg.HTTP, logger.Log(), r.Init())
	go srv.Start()

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
