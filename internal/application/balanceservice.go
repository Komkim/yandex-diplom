package application

import (
	"yandex-diplom/internal/domain/entity"
	storage "yandex-diplom/storage/repository"
)

type BalanceService struct {
	Balance *entity.Balance
}

func NewBalanceService() *BalanceService {
	return &BalanceService{}
}

func (b *BalanceService) GetBalance() (storage.Balance, error) {
	return nil, nil
}

func (b *BalanceService) SetBalanceWithdraw(order storage.Orders) error {
	return nil
}
