package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"os"
	"time"

	"os/signal"
	"syscall"
	"yandex-diplom/config"
	"yandex-diplom/internal/application"
	"yandex-diplom/internal/infrastructure/server"
	router "yandex-diplom/internal/infrastructure/server/http"

	//_ "github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
	_ "yandex-diplom/storage/postgres/migrations"
)

func main() {
	fmt.Println("start")
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		//Int("pid", os.Getpid()).
		//Str("go_version", buildInfo.GoVersion).
		Logger()
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
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

	service := application.NewServices(ctx, &cfg.Server, logger.Log())
	r := router.NewRouter(&cfg.Server, service)
	srv := server.NewServer(&cfg.HTTP, logger.Log(), r.Init())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.GetServer().Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("Server forced to shutdown:")
	}

	defer cencel()
}

func startMigration(dsn string) error {
	//dsn := "postgres://postgres:changeme@localhost:5432/yandex?sslmode=disable"
	//
	//err := connectDb(dsn)
	//if err != nil {
	//	log.Println(err)
	//}

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
