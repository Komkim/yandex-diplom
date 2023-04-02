package accrualserver

import (
	"context"
	"github.com/rs/zerolog"
	"net/http"
	"yandex-diplom/config"
)

type Server struct {
	httpServer *http.Server
	log        *zerolog.Event
}

func NewServer(cfg *config.AccrualConfig, log *zerolog.Event, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.Address,
			Handler: handler,
		},
		log: log,
	}
}

func (s *Server) Start() {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		s.log.Err(err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.log.Err(err)
	}
}

func (s *Server) GetServer() *http.Server {
	return s.httpServer
}
