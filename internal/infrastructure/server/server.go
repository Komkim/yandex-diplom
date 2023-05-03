package server

import (
	"context"
	"net/http"
	"yandex-diplom/config"

	"github.com/rs/zerolog"
)

type Server struct {
	httpServer *http.Server
	log        *zerolog.Event
}

func NewServer(cfg *config.HTTP, log *zerolog.Event, handler http.Handler) *Server {
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
