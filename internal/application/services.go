package application

import (
	"context"
	"github.com/google/uuid"
	"yandex-diplom/internal/domain/aggregate"
	"yandex-diplom/internal/mistake"

	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"yandex-diplom/config"
)

type Services struct {
	BalanceService
	OrdersService
	UsersService
	log *zerolog.Event
}

func NewServices(ctx context.Context, cfg *config.Server, log *zerolog.Event) *Services {
	db, err := newDb(ctx, cfg.DatabaseDSN)
	if err != nil {
		log.Err(err)
	}
	us := NewUsersService(db)
	os := NewOrdersService(db, us.UsersRepo)
	return &Services{
		BalanceService: NewBalanceService(db, us.UsersRepo, os.OrderRepo),
		OrdersService:  os,
		UsersService:   us,
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

func getUserID(UserRepo aggregate.UsersRepo, login string) (*uuid.UUID, error) {
	user, err := UserRepo.GetOne(login)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, mistake.UserNullError
	}
	return &user.Id, nil
}
