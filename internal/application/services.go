package application

import (
	"yandex-diplom/internal/domain/aggregate"
	"yandex-diplom/internal/mistake"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Services struct {
	BalanceService
	OrdersService
	UsersService
	log *zerolog.Event
}

func NewServices(log *zerolog.Event, db *pgxpool.Pool) *Services {
	us := NewUsersService(db)
	os := NewOrdersService(db, us.UsersRepo)
	return &Services{
		BalanceService: NewBalanceService(db, us.UsersRepo, os.OrderRepo),
		OrdersService:  os,
		UsersService:   us,
		log:            log,
	}
}

func getUserID(UserRepo aggregate.UsersRepo, login string) (*uuid.UUID, error) {
	user, err := UserRepo.GetOne(login)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, mistake.ErrUserNullError
	}
	return &user.ID, nil
}
