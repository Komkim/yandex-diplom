package storage

type Balance interface {
	GetBalance(login string) (*BalanceCurrent, error)
	SetBalanceWithdraw(withdraw *BalanceWithdraw, login string) error
	GetBalanceWithdraw(login string) (*BalanceWithdrawals, error)
	SetAccrual(accrualOrders []Order, dbOrder map[int64]Order) error
}

type BalanceCurrent struct {
	Current   *float64 `json:"current"`
	Withdrawn *float64 `json:"withdrawn"`
}

type BalanceWithdraw struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}

type BalanceWithdrawals struct {
	Order       string  `json:"order"`
	Sum         float64 `json:"sum"`
	ProcessedAt string  `json:"processed_at,omitempty"`
}
