package router

import (
	"bytes"
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"time"
	"yandex-diplom/internal/infrastructure/server/http/request"
	"yandex-diplom/internal/infrastructure/server/http/response"
	"yandex-diplom/internal/mistake"
	"yandex-diplom/pkg/luna"
)

const CookieLiveHours = 60

func (t *Router) UserRegister(w http.ResponseWriter, r *http.Request) {
	data := &request.User{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ErrInvalidRequestFormat(err))
		return
	}

	err := t.storage.Register(data.Login, data.Password)
	if errors.Is(err, mistake.LoginIsTaken) {
		render.Render(w, r, response.ErrLoginIsTaken)
		return
	}

	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}

	token := t.auth.CreateAuth(data.Login, data.Password)

	cookie := http.Cookie{
		Name:    "token",
		Path:    "/",
		Value:   token,
		Expires: time.Now().Add(time.Hour * CookieLiveHours),
	}
	http.SetCookie(w, &cookie)

	render.Render(w, r, response.UserRegistered)
}

func (t *Router) UserAuthentication(w http.ResponseWriter, r *http.Request) {
	data := &request.User{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ErrInvalidRequestFormat(err))
		return
	}

	err := t.storage.Login(data.Login, data.Password)
	if errors.Is(err, mistake.NotAuthenticated) {
		render.Render(w, r, response.ErrNotAuthenticated)
		return
	}

	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}

	token := t.auth.CreateAuth(data.Login, data.Password)

	render.Render(w, r, response.UserAuthenticated)

	cookie := http.Cookie{
		Name:    "token",
		Path:    "/",
		Value:   token,
		Expires: time.Now().Add(time.Hour * CookieLiveHours),
	}
	http.SetCookie(w, &cookie)
}

func (t *Router) OrderLoading(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		render.Render(w, r, response.ErrInvalidRequestFormat(err))
		return
	}

	number, err := strconv.ParseInt(buf.String(), 10, 64)
	if err != nil {
		render.Render(w, r, response.ErrInvalidRequestFormat(err))
		return
	}

	if !luna.Valid(number) {
		render.Render(w, r, response.OrderInvalidNumber)
		return
	}

	token, err := r.Cookie("token")
	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}

	login, _ := t.auth.FetchAuth(token.Value)

	err = t.storage.SetOrderNumber(number, login)
	if errors.Is(err, mistake.OrderAlreadyUploadedThisUser) {
		render.Render(w, r, response.OrderAlreadyBeenUploadedThis)
		return
	}
	if errors.Is(err, mistake.OrderAlreadyUploadedAnotherUser) {
		render.Render(w, r, response.OrderUploadedAnotherUser)
		return
	}

	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}

	render.Render(w, r, response.OrderNumberAccepted)
}

func (t *Router) OrderGetting(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}

	login, _ := t.auth.FetchAuth(token.Value)

	orders, err := t.storage.GetOrders(login)
	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}
	if orders == nil {
		render.Render(w, r, response.NoDataToAsnwer)
		return
	}

	//render.Render(w, r, NewOrderResponse(&orders[0]))
	orderResponse := make([]*response.OrderResponse, 0, len(orders))
	for _, o := range orders {
		temp := response.OrderResponse{
			Number:     o.Number,
			Status:     o.Status,
			Sum:        o.Accrual,
			UploadedAt: o.UploadedAt,
		}
		orderResponse = append(orderResponse, &temp)
	}
	render.RenderList(w, r, response.NewOrderListResponse(orderResponse))

	//if err := render.RenderList(w, r, NewOrderListResponse(orders)); err != nil {
	//	render.Render(w, r, ErrInternalServer(err))
	//	return
	//}

	//render.Render(w, r, SuccessfulRequestProcessing)
}

func (t *Router) BalanceCurrent(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}

	login, _ := t.auth.FetchAuth(token.Value)

	balance, err := t.storage.GetBalance(login)
	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}

	render.Render(w, r, response.BalanceToBalanceResponse(balance))
}

func (t *Router) WithdrawFounds(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}
	login, _ := t.auth.FetchAuth(token.Value)

	data := &request.OrderSumRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}

	if !luna.Valid(data.Order) {
		render.Render(w, r, response.OrderInvalidNumber)
		return
	}

	err = t.storage.SetBalanceWithdraw(&data.BalanceWithdraw, login)
	if err == mistake.BalanceNotEnouhgFunds {
		render.Render(w, r, response.BalanceNotEnouhgFunds)
		return
	}
	render.Render(w, r, response.SuccessfulRequestProcessing)
}

func (t *Router) WithdrawInformation(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}
	login, _ := t.auth.FetchAuth(token.Value)

	withdrawals, err := t.storage.GetOrderWithdrawals(login)
	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		return
	}
	if len(withdrawals) == 0 {
		render.Render(w, r, response.ErrThereIsNoWriteOff)
		return
	}

	or := make([]*response.OrderSumResponse, 0, len(withdrawals))
	for _, w := range withdrawals {
		temp := response.OrderSumResponse{
			Order:       w.Order,
			Sum:         w.Sum,
			ProcessedAt: w.ProcessedAt,
		}
		or = append(or, &temp)
	}

	render.RenderList(w, r, response.NewOrderSumListResponse(or))
}

func (t *Router) PointsAccrualsInformation(w http.ResponseWriter, r *http.Request) {

}
