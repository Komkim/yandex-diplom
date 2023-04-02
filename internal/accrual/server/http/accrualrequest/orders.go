package accrualrequest

import (
	"net/http"
	storage "yandex-diplom/storage/repository"
)

type AccrualOrdersSetRequest struct {
	storage.AccrualOrdersSet
}

func (aos *AccrualOrdersSetRequest) Bind(r *http.Request) error {
	return nil
}

type AccrualOrdersRequest struct {
	storage.AccrualOrders
}

func (aosr *AccrualOrdersRequest) Bind(r *http.Request) error {
	return nil
}
