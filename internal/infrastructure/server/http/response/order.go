package response

import (
	"github.com/go-chi/render"
	"net/http"
)

type OrderResponse struct {
	//orders storage.Order
	Number     string   `json:"number"`
	Status     string   `json:"status"`
	Sum        *float64 `json:"sum,omitempty"`
	UploadedAt string   `json:"uploaded_at"`
}

//func NewOrderResponse(order storage.Order) *OrderResponse{
//	return &OrderResponse{orders: order}
//}
//
//func NewOrderListResponse(orders []storage.Order) []render.Renderer{
//	list := []render.Renderer{}
//	for _, order := range orders {
//		list = append(list, NewOrderResponse(order))
//	}
//	return list
//}

func NewOrderListResponse(orders []*OrderResponse) []render.Renderer {
	list := []render.Renderer{}
	for _, order := range orders {
		list = append(list, order)
	}
	return list
}

func (o OrderResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type OrderSumResponse struct {
	Order       string  `json:"order"`
	Sum         float64 `json:"sum,omitempty"`
	ProcessedAt string  `json:"processed_at"`
}

func NewOrderSumListResponse(orders []*OrderSumResponse) []render.Renderer {
	list := []render.Renderer{}
	for _, order := range orders {
		list = append(list, order)
	}
	return list
}

func (o OrderSumResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
