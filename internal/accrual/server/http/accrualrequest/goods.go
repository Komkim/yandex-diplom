package accrualrequest

import (
	"net/http"
	storage "yandex-diplom/storage/repository"
)

type RewardRequest struct {
	storage.Reward
}

func (gr *RewardRequest) Bind(r *http.Request) error {
	return nil
}

type GoodsRequest struct {
	storage.Goods
}

func (gpr *GoodsRequest) Bind(r *http.Request) error {
	return nil
}
