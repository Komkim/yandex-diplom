package response

import (
	"net/http"
	storage "yandex-diplom/storage/repository"
)

type BalanceCurrentResponse struct {
	Current  *float64 `json:"current"`
	Withdraw *float64 `json:"withdraw"`
}

func (bc *BalanceCurrentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func BalanceToBalanceResponse(balance *storage.BalanceCurrent) *BalanceCurrentResponse {
	return &BalanceCurrentResponse{
		Current:  balance.Current,
		Withdraw: balance.Withdrawn,
	}
}
