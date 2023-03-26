package request

import (
	"errors"
	"net/http"
	"yandex-diplom/pkg/luna"
)

type OrderNumber struct {
	Number int64 `json:"number"`
}

func (o *OrderNumber) Bind(r *http.Request) error {
	if !luna.Valid(o.Number) {
		return errors.New("check luna algorithm false")
	}
	return nil
}
