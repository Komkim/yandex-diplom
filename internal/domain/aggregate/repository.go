package aggregate

import (
	"github.com/google/uuid"
	"yandex-diplom/internal/domain/entity"
)

const DBTIMEOUT = 5

type UsersRepo interface {
	GetOne(login string) (*entity.Users, error)
	SetOne(login, password string) error
}

type OrdersRepo interface {
	SetOne(number int64, userID uuid.UUID) error
	SetSum(number int64, userID uuid.UUID, sum float64) error
	GetAllByUser(userID uuid.UUID) ([]entity.Orders, error)
	GetAllByUserWithdrawals(userid uuid.UUID) ([]entity.Orders, error)
	GetByNumber(number int64) (*entity.Orders, error)
}

type BalanceRepo interface {
	GetByUser(userID uuid.UUID) (*entity.Balance, error)
	SetSum(userID uuid.UUID, sum float64) error
	GetCurrentByUser(userId uuid.UUID) (*float64, error)
	GetWithdrawntByUser(userId uuid.UUID) (*float64, error)
}
