package application

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"yandex-diplom/internal/domain/aggregate"
	storage "yandex-diplom/storage/repository"
)

type OrdersService struct {
	OrderRepo *aggregate.OrdersRepo
}

func NewOrdersService(db *pgxpool.Pool) OrdersService {
	or := aggregate.NewOrdersRepo(db)
	return OrdersService{OrderRepo: or}
}

func (o *OrdersService) SetOrderNumber(number int64) error {
	//ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	//defer cancel()

	return nil
}
func (o *OrdersService) GetOrders() ([]storage.Orders, error) {
	return nil, nil
}
func (o *OrdersService) GetOrderWithdrawals() ([]storage.Orders, error) {
	return nil, nil
}
