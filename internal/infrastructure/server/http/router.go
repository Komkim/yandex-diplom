package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
	"yandex-diplom/config"
	"yandex-diplom/internal/infrastructure/server/auth"
	storage "yandex-diplom/storage/repository"
)

type Router struct {
	cfg     *config.Server
	storage storage.Storage
	auth    auth.AuthInterface
}

func NewRouter(cfg *config.Server, storage storage.Storage, auth auth.AuthInterface) *Router {
	return &Router{
		cfg:     cfg,
		storage: storage,
		auth:    auth,
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

		r.Post("/register", t.UserRegister)
		r.Post("/login", t.UserAuthentication)

		r.Group(func(r chi.Router) {

			r.Use(t.AuthMiddleware)

			r.Post("/orders", t.OrderLoading)
			r.Get("/orders", t.OrderGetting)

			r.Get("/balance", t.BalanceCurrent)
			r.Post("/balance/withdraw", t.WithdrawFounds)
			r.Get("/withdraw", t.WithdrawInformation)

			r.Get("/orders/{orderId}", t.PointsAccrualsInformation)
		})
	})

	return router
}
