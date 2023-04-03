package accrualrouter

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"yandex-diplom/internal/accrual/server/http/accrualrequest"
	"yandex-diplom/internal/accrual/server/http/accrualresponse"
	"yandex-diplom/pkg/luna"
)

func (t *AccrualRouter) GoodsLoading(w http.ResponseWriter, r *http.Request) {
	data := &accrualrequest.RewardRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, accrualresponse.ErrInvalidRequestFormat(err))
		t.log.Err(err)
		return
	}

	err := t.storage.SetReward(&data.Reward)
	if err != nil {
		render.Render(w, r, accrualresponse.ErrInternalServer(err))
		t.log.Err(err)
		return
	}

	render.Render(w, r, accrualresponse.RewardRegistered)
}

func (t *AccrualRouter) OrdersLoading(w http.ResponseWriter, r *http.Request) {
	data := &accrualrequest.AccrualOrdersSetRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, accrualresponse.ErrInvalidRequestFormat(err))
		t.log.Err(err)
		return
	}

	n, err := strconv.ParseInt(data.Order, 10, 64)
	if err != nil {
		render.Render(w, r, accrualresponse.ErrInvalidRequestFormat(err))
		t.log.Err(err)
		return
	}
	if !luna.Valid(n) {
		render.Render(w, r, accrualresponse.ErrInvalidRequestFormat(errors.New("luna not valid")))
		t.log.Err(errors.New("luna not valid"))
		return
	}

	err = t.storage.SetAccrualOrders(&data.AccrualOrdersSet)
	if err != nil {
		render.Render(w, r, accrualresponse.ErrInternalServer(err))
		t.log.Err(err)
		return
	}

	render.Render(w, r, accrualresponse.OrderSuccessfully)

}

func (t *AccrualRouter) OrdersInformation(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.ParseInt(chi.URLParam(r, "number"), 0, 64)
	if err != nil {
		render.Render(w, r, accrualresponse.ErrInternalServer(err))
		t.log.Err(err)
		return
	}

	//number := r.Context().Value("number").(*int64)

	order, err := t.storage.GetOrderByNumber(number)
	if err != nil {
		render.Render(w, r, accrualresponse.ErrInternalServer(err))
		t.log.Err(err)
		return
	}

	if order == nil {
		render.Render(w, r, accrualresponse.OrderNotRegistered)
		return
	}

	render.Render(w, r, accrualresponse.OrderToResponse(order))
}
