package application

import (
	"context"
	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"yandex-diplom/config"
)

type Services struct {
	BalanceService
	OrdersService
	UsersService
	DbPool *pgxpool.Pool
	log    *zerolog.Event
}

func NewServices(ctx context.Context, cfg *config.Server, log *zerolog.Event) *Services {
	db, err := newDb(ctx, cfg.DatabaseDSN)
	if err != nil {
		log.Err(err)
	}
	return &Services{
		BalanceService: NewBalanceService(),
		OrdersService:  NewOrdersService(),
		UsersService:   NewUsersService(),
		DbPool:         db,
		log:            log,
	}
}

func newDb(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
