package server

import (
	"errors"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(address string, router http.Handler, readTO, writeTO, idleTO time.Duration) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         address,
			Handler:      router,
			ReadTimeout:  readTO,
			WriteTimeout: writeTO,
			IdleTimeout:  idleTO,
		},
	}
}

func (s *Server) Run() error {
	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
