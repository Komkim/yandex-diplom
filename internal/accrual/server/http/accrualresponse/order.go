package accrualresponse

import (
	"net/http"
	storage "yandex-diplom/storage/repository"
)

type OrderResponse struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

func OrderToResponse(orders *storage.AccrualOrders) *OrderResponse {
	return &OrderResponse{
		Order:   orders.Order,
		Status:  orders.Status,
		Accrual: orders.Accrual,
	}
}
func (o OrderResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
