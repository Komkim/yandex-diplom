package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
	"yandex-diplom/config"
	storage "yandex-diplom/storage/repository"
)

type Router struct {
	cfg     *config.Server
	storage storage.Storage
}

func NewRouter(cfg *config.Server, storage storage.Storage) *Router {
	return &Router{
		cfg:     cfg,
		storage: storage,
	}
}

func (t *Router) Init() http.Handler {
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

		router.Post("/register", t.UserRegister)
		router.Post("/login", t.UserAuthentication)

		router.Post("/orders", t.OrderLoading)
		router.Get("/orders", t.OrderGetting)

		router.Get("/balance", t.BalanceCurrent)
		router.Post("/balance/withdraw", t.WithdrawFounds)
		router.Get("/withdraw", t.WithdrawInformation)

		router.Get("/orders/{orderId}", t.PointsAccrualsInformation)
	})

	return router
}
