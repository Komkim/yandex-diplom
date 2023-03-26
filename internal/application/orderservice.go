package application

import (
	"yandex-diplom/internal/domain/entity"
	storage "yandex-diplom/storage/repository"
)

type OrdersService struct {
	Orders entity.Orders
}

func NewOrdersService() OrdersService {
	return OrdersService{}
}

func (o *OrdersService) SetOrderNumber(number int64) error {
	return nil
}
func (o *OrdersService) GetOrders() ([]storage.Orders, error) {
	return nil, nil
}
func (o *OrdersService) GetOrderWithdrawals() ([]storage.Orders, error) {
	return nil, nil
}
