package router

import (
	"bytes"
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"time"
	"yandex-diplom/internal/infrastructure/server/http/request"
	"yandex-diplom/internal/mistake"
	"yandex-diplom/pkg/luna"
)

const CookieLiveHours = 60

func (t *Router) UserRegister(w http.ResponseWriter, r *http.Request) {
	data := &request.User{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequestFormat(err))
		return
	}

	err := t.storage.Register(data.Login, data.Password)
	if errors.Is(err, mistake.LoginIsTaken) {
		render.Render(w, r, ErrLoginIsTaken)
		return
	}

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
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

	render.Render(w, r, UserRegistered)
}

func (t *Router) UserAuthentication(w http.ResponseWriter, r *http.Request) {
	data := &request.User{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequestFormat(err))
		return
	}

	err := t.storage.Login(data.Login, data.Password)
	if errors.Is(err, mistake.NotAuthenticated) {
		render.Render(w, r, ErrNotAuthenticated)
		return
	}

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	token := t.auth.CreateAuth(data.Login, data.Password)

	render.Render(w, r, UserAuthenticated)

	cookie := http.Cookie{
		Name:    "token",
		Path:    "/",
		Value:   token,
		Expires: time.Now().Add(time.Hour * CookieLiveHours),
	}
	http.SetCookie(w, &cookie)
}

func (t *Router) OrderLoading(w http.ResponseWriter, r *http.Request) {
	login, pass, ok := r.BasicAuth()
	if ok {
		render.Render(w, r, ErrNotAuthenticated)
		return
	}
	err := t.storage.Login(login, pass)
	if errors.Is(err, mistake.NotAuthenticated) {
		render.Render(w, r, ErrNotAuthenticated)
		return
	}

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		render.Render(w, r, ErrInvalidRequestFormat(err))
		return
	}

	number, err := strconv.ParseInt(buf.String(), 10, 64)
	if err != nil {
		render.Render(w, r, ErrInvalidRequestFormat(err))
		return
	}

	if luna.Valid(number) {
		render.Render(w, r, ErrInvalidOrderNumber)
		return
	}

	err = t.storage.SetOrderNumber(number)
	if err != nil {
		render.Render(w, r, ErrInvalidRequestFormat(err))
		return
	}

	render.Render(w, r, OrderAlreadyBeenUploaded)
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
