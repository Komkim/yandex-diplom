package application

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"yandex-diplom/internal/domain/aggregate"
	"yandex-diplom/internal/domain/valueobject"
	"yandex-diplom/internal/mistake"
	storage "yandex-diplom/storage/repository"
)

type BalanceService struct {
	BalanceRepo aggregate.BalanceRepo
	UserRepo    aggregate.UsersRepo
	OrderRepo   aggregate.OrdersRepo
}

func NewBalanceService(db *pgxpool.Pool, userRepo aggregate.UsersRepo, orderRepo aggregate.OrdersRepo) BalanceService {
	return BalanceService{
		BalanceRepo: aggregate.NewBalanceRepo(db),
		UserRepo:    userRepo,
		OrderRepo:   orderRepo,
	}
}

func (b *BalanceService) GetBalance(login string) (*storage.BalanceCurrent, error) {
	userID, err := getUserID(b.UserRepo, login)
	if err != nil {
		return nil, err
	}

	current, err := b.BalanceRepo.GetCurrentByUser(*userID)
	if err != nil {
		return nil, err
	}

	withdraw, err := b.BalanceRepo.GetWithdrawntByUser(*userID)
	if err != nil {
		return nil, err
	}

	return &storage.BalanceCurrent{
		Current:   current,
		Withdrawn: withdraw,
	}, nil
}

func (b *BalanceService) SetBalanceWithdraw(withdraw *storage.BalanceWithdraw, login string) error {
	userID, err := getUserID(b.UserRepo, login)
	if err != nil {
		return err
	}

	current, err := b.BalanceRepo.GetCurrentByUser(*userID)
	if err != nil {
		return err
	}
	if current != nil && *current < withdraw.Sum {
		return mistake.ErrBalanceNotEnouhgFunds
	}

	w := -withdraw.Sum
	err = b.OrderRepo.SetSum(withdraw.Order, *userID, w)
	if err != nil {
		return err
	}
	err = b.BalanceRepo.SetSum(*userID, w)
	if err != nil {
		return err
	}

	return nil
}

func (b *BalanceService) GetBalanceWithdraw(login string) (*storage.BalanceWithdrawals, error) {
	return nil, nil
}

func (b *BalanceService) SetAccrual(accrualOrders []storage.Order, dbOrder map[int64]storage.Order) error {
f:
	for _, or := range accrualOrders {
		if or.Status == dbOrder[or.Number].Status {
			continue f
		}

		entityOrder, err := b.OrderRepo.GetByNumber(or.Number)
		if err != nil {
			return err
		}

		err = b.OrderRepo.SetOrder(or.Number, entityOrder.UserID, *or.Accrual, or.Status)
		if err != nil {
			return err
		}

		if or.Status == valueobject.PROCESSED {
			err = b.BalanceRepo.SetSum(entityOrder.UserID, *or.Accrual)
			if err != nil {
				return err
			}
		}
	}

	//Получаем из базы по номерам все ордера что нам передали. Надо чтобы получить юзеров
	//Сравниваем все полученные ордера с теми что прислал акруал. Если есть изменения записываем их в ордера
	//Для ордеров со статуом завершено делаем запись для баланса

	return nil
}
