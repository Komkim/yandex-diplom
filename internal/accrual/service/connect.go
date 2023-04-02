package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"yandex-diplom/internal/accrual/model"
	storage "yandex-diplom/storage/repository"
)

type ConnectService struct {
	ConnectRepo model.ConnectRepo
}

func NewConnectService(db *pgxpool.Pool) ConnectService {
	return ConnectService{
		ConnectRepo: model.NewConnect(db),
	}
}

func (cs *ConnectService) SetConnect(goods *storage.Goods, order *storage.AccrualOrders) error {
	return nil
}
