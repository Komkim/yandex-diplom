package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
	"yandex-diplom/internal/accrual/model"
	storage "yandex-diplom/storage/repository"
)

type AccrualOrdersService struct {
	OrdersRepo  model.OrdersRepo
	GoodsRepo   model.GoodsRepo
	ConnectRepo model.ConnectRepo
}

func NewAccrualOrdersService(db *pgxpool.Pool, gr model.GoodsRepo, cr model.ConnectRepo) AccrualOrdersService {
	return AccrualOrdersService{
		OrdersRepo:  model.NewOrders(db),
		GoodsRepo:   gr,
		ConnectRepo: cr,
	}
}

func (o *AccrualOrdersService) SetAccrualOrders(ordesSet *storage.AccrualOrdersSet) error {
	var r float64
	n, err := strconv.ParseInt(ordesSet.Order, 10, 64)
	if err != nil {
		return err
	}

	for _, good := range ordesSet.Goods {
		reward, err := o.GoodsRepo.GetReward(good.Description)
		if err != nil {
			return err
		}

		r += reward.Reward

		goodsID, err := o.GoodsRepo.SetGoods(good.Description, good.Price)
		if err != nil {
			return err
		}

		err = o.ConnectRepo.SetConnect(n, *goodsID)
		if err != nil {
			return err
		}
	}

	err = o.OrdersRepo.SetOrders(n, model.PROCESSED, r)
	if err != nil {
		return err
	}

	return nil
}

func (o *AccrualOrdersService) GetOrderByNumber(number int64) (*storage.AccrualOrders, error) {
	orders, err := o.OrdersRepo.GetOrdersByNumber(number)
	if err != nil {
		return nil, err
	}
	if orders == nil {
		return nil, nil
	}

	return &storage.AccrualOrders{
		Order:   strconv.FormatInt(orders.Number, 10),
		Status:  orders.Status,
		Accrual: orders.Reward,
	}, nil
}
