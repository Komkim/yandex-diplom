package request

import (
	"net/http"
	storage "yandex-diplom/storage/repository"
)

type OrderSumRequest struct {
	//Order string  `json:"order"`
	//Sum   float64 `json:"sum"`
	storage.BalanceWithdraw
}

func (os *OrderSumRequest) Bind(r *http.Request) error {
	return nil
}
