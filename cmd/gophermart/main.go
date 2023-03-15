package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"yandex-diplom/config"
	"yandex-diplom/internal/application"
	"yandex-diplom/internal/infrastructure/server"
	router "yandex-diplom/internal/infrastructure/server/http"
	storage "yandex-diplom/storage/repository"

	//_ "github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
	_ "yandex-diplom/storage/postgres/migrations"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Debug().Str("server", "start")

	ctx, cencel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		log.Error().Err(err)
	}

	err = startMigration(cfg.Server.DatabaseDSN)
	if err != nil {
		log.Error().Err(err)
	}

	service := application.NewServices(ctx, &cfg.Server, log.Log())
	r := router.NewRouter(&cfg.Server, service)
	srv := server.NewServer(&cfg.HTTP, log.Log(), r.Init())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.GetServer().Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server forced to shutdown:")
	}

	defer cencel()
	//

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
