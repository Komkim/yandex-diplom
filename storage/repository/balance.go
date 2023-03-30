package storage

type Balance interface {
	GetBalance(login string) (*BalanceCurrent, error)
	SetBalanceWithdraw(withdraw *BalanceWithdraw, login string) error
	GetBalanceWithdraw(login string) (*BalanceWithdrawals, error)
}

type BalanceCurrent struct {
	Current   *float64 `json:"current"`
	Withdrawn *float64 `json:"withdrawn"`
}

type BalanceWithdraw struct {
	Order int64   `json:"order"`
	Sum   float64 `json:"sum"`
}

type BalanceWithdrawals struct {
	Order       int64   `json:"order"`
	Sum         float64 `json:"sum"`
	ProcessedAt string  `json:"processed_at,omitempty"`
}
