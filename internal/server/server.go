package server

import (
	"context"
	"net/http"
)

type Server struct {
	mux    *http.ServeMux
	server *http.Server
}

func NewServer(handler *http.ServeMux) *Server {
	return &Server{
		mux: handler,
		server: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
