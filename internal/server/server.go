package server

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"net/http"
)

type Server struct {
	e *echo.Echo // Використовуємо echo.Echo безпосередньо
}

func NewServer(handler *echo.Echo) *Server {
	return &Server{e: handler} // Призначаємо наданий екземпляр Echo
}

func (s *Server) Run() error {
	return s.e.StartServer(&http.Server{Addr: ":8080"}) // Запускаємо сервер Echo
}

func (s *Server) Stop(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
