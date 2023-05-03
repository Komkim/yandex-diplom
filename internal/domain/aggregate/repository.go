package aggregate

import (
	"yandex-diplom/internal/domain/entity"

	"github.com/google/uuid"
)

const DBTIMEOUT = 5

type UsersRepo interface {
	GetOne(login string) (*entity.Users, error)
	SetOne(login, password string) error
}

type OrdersRepo interface {
	SetOne(number int64, userID uuid.UUID) error
	SetSum(number int64, userID uuid.UUID, sum float64) error
	SetOrder(number int64, userID uuid.UUID, sum float64, status string) error
	GetAllByUser(userID uuid.UUID) ([]entity.Orders, error)
	GetAllByUserWithdrawals(userID uuid.UUID) ([]entity.Orders, error)
	GetByNumber(number int64) (*entity.Orders, error)
	GetAccrualPoll() ([]entity.Orders, error)
}

type BalanceRepo interface {
	GetByUser(userID uuid.UUID) (*entity.Balance, error)
	SetSum(userID uuid.UUID, sum float64) error
	GetCurrentByUser(userID uuid.UUID) (*float64, error)
	GetWithdrawntByUser(userID uuid.UUID) (*float64, error)
}
