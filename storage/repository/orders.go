package storage

type Orders interface {
	SetOrderNumber(number int64, login string) error
	//GetOrderByNymber(number int64, login string) (*Order, error)
	GetOrders(login string) ([]Order, error)
	GetOrderWithdrawals(login string) ([]OrderWithdrawals, error)
}

type Order struct {
	Number     int64    `json:"number"`
	Status     string   `json:"status"`
	Accrual    *float64 `json:"accrual,omitempty"`
	UploadedAt string   `json:"uploaded_at"`
}

type OrderWithdrawals struct {
	Order       int64   `json:"order"`
	Sum         float64 `json:"sum"`
	ProcessedAt string  `json:"processed_at,omitempty"`
}
