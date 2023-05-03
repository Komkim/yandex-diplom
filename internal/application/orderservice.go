package application

import (
	"strconv"
	"time"
	"yandex-diplom/internal/domain/aggregate"
	"yandex-diplom/internal/domain/entity"
	"yandex-diplom/internal/mistake"
	storage "yandex-diplom/storage/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrdersService struct {
	OrderRepo aggregate.OrdersRepo
	UserRepo  aggregate.UsersRepo
}

func NewOrdersService(db *pgxpool.Pool, userRepo aggregate.UsersRepo) OrdersService {
	return OrdersService{
		OrderRepo: aggregate.NewOrdersRepo(db),
		UserRepo:  userRepo,
	}
}

func (o *OrdersService) SetOrderNumber(number int64, login string) error {
	userID, err := getUserID(o.UserRepo, login)
	if err != nil {
		return err
	}

	order, err := o.OrderRepo.GetByNumber(number)
	if err != nil {
		return err
	}
	if order != nil {
		if order.UserID == *userID {
			return mistake.ErrOrderAlreadyUploadedThisUser
		}
		if order.UserID != *userID {
			return mistake.ErrOrderAlreadyUploadedAnotherUser
		}
	}

	return o.OrderRepo.SetOne(number, *userID)

}

func (o *OrdersService) GetOrders(login string) ([]storage.Order, error) {
	userID, err := getUserID(o.UserRepo, login)
	if err != nil {
		return nil, err
	}

	entityOrders, err := o.OrderRepo.GetAllByUser(*userID)
	if err != nil {
		return nil, err
	}

	orders := converOrders(entityOrders)

	return orders, nil
}
func (o *OrdersService) GetOrderWithdrawals(login string) ([]storage.OrderWithdrawals, error) {
	userID, err := getUserID(o.UserRepo, login)
	if err != nil {
		return nil, err
	}
	orders, err := o.OrderRepo.GetAllByUserWithdrawals(*userID)
	if err != nil {
		return nil, err
	}

	ow := make([]storage.OrderWithdrawals, 0, len(orders))
	for _, v := range orders {
		var s float64
		if v.Sum != nil {
			s = *v.Sum
		}
		ow = append(
			ow,
			storage.OrderWithdrawals{
				Order:       strconv.FormatInt(v.Number, 10),
				Sum:         -s,
				ProcessedAt: v.CreateAt.Format(time.RFC3339),
			})
	}

	return ow, nil
}

func converOrders(entityOrders []entity.Orders) []storage.Order {
	o := make([]storage.Order, 0, len(entityOrders))
	for _, order := range entityOrders {
		c := order.CreateAt.Format(time.RFC3339)
		tempOrder := storage.Order{
			Number:     strconv.FormatInt(order.Number, 10),
			Status:     order.Status,
			Accrual:    order.Sum,
			UploadedAt: c,
		}
		o = append(o, tempOrder)
	}
	return o
}

func (o *OrdersService) GetAccrualOrder() ([]storage.Order, error) {
	entityOrders, err := o.OrderRepo.GetAccrualPoll()
	if err != nil {
		return nil, err
	}

	storageOrder := converOrders(entityOrders)

	return storageOrder, err
}
