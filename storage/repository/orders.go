package storage

type Orders interface {
	SetOrderNumber(number int64) error
	GetOrders() ([]Orders, error)
	GetOrderWithdrawals() ([]Orders, error)
}
