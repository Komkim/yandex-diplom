package model

import "github.com/google/uuid"

const (
	DBTIMEOUT = 5

	PERCENT  = "%"
	QUANTITY = "pt"

	REGISTERED = "REGISTERED"
	INVALID    = "INVALID"
	PROCESSING = "PROCESSING"
	PROCESSED  = "PROCESSED"
)

type ConnectRepo interface {
	SetConnect(number int64, goodsID uuid.UUID) error
}

type GoodsRepo interface {
	SetGoods(description string, price float64) (*uuid.UUID, error)
	GetGoods(description string) (*ModelGoods, error)
	SetReward(name, rewardType string, reward int64) error
	GetReward(description string) (*ModelReward, error)
}

type OrdersRepo interface {
	SetOrders(number int64, status string, reward float64) error
	GetOrdersByNumber(number int64) (*ModelOrders, error)
}
