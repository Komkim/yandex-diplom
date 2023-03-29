package application

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"yandex-diplom/internal/domain/aggregate"
	"yandex-diplom/internal/domain/entity"
	"yandex-diplom/internal/mistake"
	storage "yandex-diplom/storage/repository"
)

type OrdersService struct {
	OrderRepo aggregate.OrdersRepo
	UserRepo  aggregate.UsersRepo
}

func NewOrdersService(db *pgxpool.Pool, userRepo aggregate.UsersRepo) OrdersService {
	or := aggregate.NewOrdersRepo(db)
	return OrdersService{
		OrderRepo: or,
		UserRepo:  userRepo,
	}
}

func (o *OrdersService) SetOrderNumber(number int64, login string) error {
	user, err := o.UserRepo.GetOne(login)
	if err != nil {
		return err
	}
	if user == nil {
		return mistake.UserNullError
	}

	order, err := o.OrderRepo.GetByNumber(number)
	if err != nil {
		return err
	}
	if order != nil {
		if order.UserId == user.Id {
			return mistake.OrderAlreadyUploadedThisUser
		}
		if order.UserId != user.Id {
			return mistake.OrderAlreadyUploadedAnotherUser
		}
	}

	return o.OrderRepo.SetOne(number, user.Id)

}

//	func (o *OrdersService) GetOrderByNymber(number int64, login string) (*storage.Order, error) {
//		user, err := o.UserRepo.GetOne(login)
//		if err != nil {
//			return nil, err
//		}
//		if user == nil {
//			return nil, mistake.UserNullError
//		}
//
//		order, err := o.OrderRepo.GetByNumber(number, user.Id)
//		if err != nil {
//			return nil, err
//		}
//		return
//	}
func (o *OrdersService) GetOrders(login string) ([]storage.Order, error) {
	user, err := o.UserRepo.GetOne(login)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, mistake.UserNullError
	}

	entityOrders, err := o.OrderRepo.GetAllByUser(user.Id)
	if err != nil {
		return nil, err
	}

	orders, err := converOrders(entityOrders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
func (o *OrdersService) GetOrderWithdrawals() ([]storage.Order, error) {
	return nil, nil
}

func converOrders(entityOrders []entity.Orders) ([]storage.Order, error) {
	o := make([]storage.Order, 0, len(entityOrders))
	for _, order := range entityOrders {
		c := order.CreateAt.Format(time.RFC3339)
		//createAt, err := time.Parse(time.RFC3339, c)
		//createAt, err := time.Parse(time.RFC3339, order.CreateAt.String())
		//if err != nil {
		//	return nil, err
		//}
		tempOrder := storage.Order{
			Number: order.Number,
			//Status:     order.Status.String(),
			Status:     order.Status,
			Accrual:    order.Accrual,
			UploadedAt: c,
		}
		o = append(o, tempOrder)
	}
	return o, nil
}
