package request

import (
	"net/http"
	"yandex-diplom/internal/mistake"
	"yandex-diplom/pkg/luna"
)

type OrderNumber struct {
	Number int64 `json:"number"`
}

func (o *OrderNumber) Bind(r *http.Request) error {
	if !luna.Valid(o.Number) {
		return mistake.InvalidOrderNumber
	}
	return nil
}
