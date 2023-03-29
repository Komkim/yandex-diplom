package aggregate

import (
	"github.com/google/uuid"
	"yandex-diplom/internal/domain/entity"
)

type UsersRepo interface {
	GetOne(login string) (*entity.Users, error)
	SetOne(login, password string) error
}

type OrdersRepo interface {
	SetOne(number int64, userId uuid.UUID) error
	GetAllByUser(userid uuid.UUID) ([]entity.Orders, error)
	GetAllByUserWithdrawals(userid string) ([]entity.Orders, error)
	GetByNumber(number int64) (*entity.Orders, error)
}
