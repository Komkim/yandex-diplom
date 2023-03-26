package router

import (
	"github.com/go-chi/render"
	"net/http"
	"yandex-diplom/internal/infrastructure/server/http/request"
)

//type ErrorResponse struct {
//	Err        error  `json:"-"`
//	StatusCode int    `json:"-"`
//	StatusText string `json:"status_text"`
//	Message    string `json:"message"`
//}
//
//var (
//	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
//	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
//	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
//)

func (t *Router) UserRegister(w http.ResponseWriter, r *http.Request) {
	//t.storage.Register()
}

func (t *Router) UserAuthentication(w http.ResponseWriter, r *http.Request) {

}

func (t *Router) OrderLoading(w http.ResponseWriter, r *http.Request) {

	n := &request.OrderNumber{}
	if err := render.Bind(r, n); err != nil {
		render.Render(w, r, ErrInvalidRequestFormat(err))
		return
	}

	err := t.storage.SetOrderNumber(n.Number)
	if err != nil {
		render.Render(w, r, ErrInvalidRequestFormat(err))
		return
	}

	render.Status(r, http.StatusCreated)
	//render.Render(w, r, http.StatusOK)
}

func (t *Router) OrderGetting(w http.ResponseWriter, r *http.Request) {

}

func (t *Router) BalanceCurrent(w http.ResponseWriter, r *http.Request) {

}

func (t *Router) WithdrawFounds(w http.ResponseWriter, r *http.Request) {

}

func (t *Router) WithdrawInformation(w http.ResponseWriter, r *http.Request) {

}

func (t *Router) PointsAccrualsInformation(w http.ResponseWriter, r *http.Request) {

}
