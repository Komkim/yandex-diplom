package storage

type Balance interface {
	GetBalance() (Balance, error)
	SetBalanceWithdraw(order Orders) error
}
