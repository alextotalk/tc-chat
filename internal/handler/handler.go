package handler

import (
	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/alextotalk/tc-chat/internal/service"
	"html/template"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Pass any necessary data to the template (optional)
	users := []domain.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	data := map[string]interface{}{
		"Users": users,
	}
	tpl.Execute(w, data) // Write to response writer
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("templates/auth.html"))
	tpl.Execute(w, nil) // Write to response writer
}

func (h *Handler) initRoutes() {
	h.HandleFunc("/", indexHandler)
	h.HandleFunc("/auth", authHandler)
}

func (h *Handler) NewRouter() *http.ServeMux {
	return h.mux
}
