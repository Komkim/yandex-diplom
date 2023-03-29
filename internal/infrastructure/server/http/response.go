package router

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequestFormat(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "invalid request format.",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServer(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error",
		ErrorText:      err.Error(),
	}
}

var UserRegistered = &Response{HTTPStatusCode: 200, StatusText: "User successfully registered and authenticated."}
var UserAuthenticated = &Response{HTTPStatusCode: 200, StatusText: "User successfully authenticated."}

var OrderAlreadyBeenUploadedThis = &Response{HTTPStatusCode: 200, StatusText: "Order number has already been uploaded by this user."}
var OrderNumberAccepted = &Response{HTTPStatusCode: 202, StatusText: "New order number accepted for processing."}
var OrderUploadedAnotherUser = &Response{HTTPStatusCode: 409, StatusText: "The order number has already been uploaded by another user."}
var OrderInvalidNumber = &Response{HTTPStatusCode: 422, StatusText: "Invalid order number."}

var SuccessfulRequestProcessing = &Response{HTTPStatusCode: 200, StatusText: "Successful request processing."}
var NoDataToAsnwer = &Response{HTTPStatusCode: 204, StatusText: "No data to answer."}

var ErrNoDataToAnswer = &Response{HTTPStatusCode: 204, StatusText: "no data to answer."}
var ErrThereIsNoWriteOff = &Response{HTTPStatusCode: 204, StatusText: "there is no write-off."}

var ErrNotAuthorized = &Response{HTTPStatusCode: 401, StatusText: "User is not authorized."}
var ErrNotAuthenticated = &Response{HTTPStatusCode: 401, StatusText: "User not authenticated."}
var ErrNamePassPair = &Response{HTTPStatusCode: 401, StatusText: "invalid username/password pair"}
var ErrNotEnoughFunds = &Response{HTTPStatusCode: 402, StatusText: "There are not enough funds on the account."}
var ErrNotFound = &Response{HTTPStatusCode: 404, StatusText: "Resource not found."}
var ErrMethodNotAllowed = &Response{HTTPStatusCode: 405, StatusText: "Method not allowed"}
var ErrLoginIsTaken = &Response{HTTPStatusCode: 409, StatusText: "Login is taken."}

type OrderResponse struct {
	//orders storage.Order
	Number     int64    `json:"number"`
	Status     string   `json:"status"`
	Accrual    *float64 `json:"accrual,omitempty"`
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
