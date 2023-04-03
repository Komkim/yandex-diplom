package application

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
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

	num, err := strconv.ParseInt(withdraw.Order, 10, 64)
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
	err = b.OrderRepo.SetSum(num, *userID, w)
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

func (b *BalanceService) SetAccrual(accrualOrders []storage.OrderAccrual) error {
f:
	for _, or := range accrualOrders {
		num, err := strconv.ParseInt(or.Order, 10, 64)
		if err != nil {
			return err
		}

		entityOrder, err := b.OrderRepo.GetByNumber(num)
		if err != nil {
			return err
		}

		if or.Status == entityOrder.Status {
			continue f
		}

		err = b.OrderRepo.SetOrder(num, entityOrder.UserID, or.Accrual, or.Status)
		if err != nil {
			return err
		}

		if or.Status == valueobject.PROCESSED {
			err = b.BalanceRepo.SetSum(entityOrder.UserID, or.Accrual)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
