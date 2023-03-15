package router

import "net/http"

func (t *Router) UserRegister(w http.ResponseWriter, r *http.Request) {
	t.storage.Register()
}

func (t *Router) UserAuthentication(w http.ResponseWriter, r *http.Request) {

}

func (t *Router) OrderLoading(w http.ResponseWriter, r *http.Request) {

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
