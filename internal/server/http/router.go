package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
	"yandex-diplom/config"
)

type Router struct {
	cfg *config.Server
}

func NewRouter(cfg *config.Server) *Router {
	return &Router{
		cfg: cfg,
	}
}

func (r *Router) Init() http.Handler {
	router := chi.NewRouter()

	//router.Use(middleware.RequestID)
	//router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))

	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Route("/api/user", func(r chi.Router) {

		router.Post("/register", UserRegister)
		router.Post("/login", UserAuthentication)

		router.Post("/orders", OrderLoading)
		router.Get("/orders", OrderGetting)

		router.Get("/balance", BalanceCurrent)
		router.Post("/balance/withdraw", WithdrawFounds)
		router.Get("/withdraw", WithdrawInformation)

		router.Get("/orders/{orderId}", PointsAccrualsInformation)
	})

	return router
}
