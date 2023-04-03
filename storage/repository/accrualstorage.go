package storage

type AccrualStorage interface {
	IGoods
	IOrdersAccrual
	IConnect
}

type IGoods interface {
	SetGoods(goods *Goods) error
	SetReward(goods *Reward) error
}

type IOrdersAccrual interface {
	SetAccrualOrders(orser *AccrualOrdersSet) error
	GetOrderByNumber(number int64) (*AccrualOrders, error)
}

type IConnect interface {
	SetConnect(goods *Goods, order *AccrualOrders) error
}

type Reward struct {
	Match      string `json:"match"`
	Reward     int64  `json:"reward"`
	RewardType string `json:"reward_type"`
}

type Goods struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type AccrualOrdersSet struct {
	Order string  `json:"order"`
	Goods []Goods `json:"goods"`
}

type AccrualOrders struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}
