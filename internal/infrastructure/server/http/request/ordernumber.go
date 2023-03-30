package request

import (
	"net/http"
	"yandex-diplom/internal/mistake"
	"yandex-diplom/pkg/luna"
	storage "yandex-diplom/storage/repository"
)

type OrderNumber struct {
	Number int64 `json:"number"`
}

func (o *OrderNumber) Bind(r *http.Request) error {
	if !luna.Valid(o.Number) {
		return mistake.OrderInvalidNumber
	}
	return nil
}

type OrderSumRequest struct {
	//Order int64 `json:"order"`
	//Sum float64 `json:"sum"`
	storage.BalanceWithdraw
}

func (os *OrderSumRequest) Bind(r *http.Request) error {
	return nil
}
