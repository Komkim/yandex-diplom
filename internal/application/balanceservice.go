package application

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"yandex-diplom/internal/domain/aggregate"
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
		return mistake.BalanceNotEnouhgFunds
	}

	w := -withdraw.Sum
	err = b.OrderRepo.SetSum(withdraw.Order, *userID, w)
	if err != nil {
		return err
	}
	err = b.BalanceRepo.SetSum(*userID, withdraw.Sum)
	if err != nil {
		return err
	}

	return nil
}

func (b *BalanceService) GetBalanceWithdraw(login string) (*storage.BalanceWithdrawals, error) {
	return nil, nil
}
