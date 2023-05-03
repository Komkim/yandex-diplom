package accrualrouter

import (
	"net/http"
	"time"
	"yandex-diplom/config"
	storage "yandex-diplom/storage/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

type AccrualRouter struct {
	cfg     *config.AccrualConfig
	storage storage.AccrualStorage
	log     *zerolog.Event
}

func NewAccrualRouter(cfg *config.AccrualConfig, storage storage.AccrualStorage, log *zerolog.Event) *AccrualRouter {
	return &AccrualRouter{
		cfg:     cfg,
		storage: storage,
		log:     log,
	}
}

func (t *AccrualRouter) Init() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))

	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Route("/api", func(r chi.Router) {

		r.Post("/orders", t.OrdersLoading)
		r.Get("/orders/{number}", t.OrdersInformation)

		r.Post("/goods", t.GoodsLoading)

	})

	return router
}
