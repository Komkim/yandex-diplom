package storage

type Orders interface {
	SetOrderNumber(number string) error
	GetOrders() ([]Orders, error)
	GetOrderWithdrawals() ([]Orders, error)
}
