package handler

import (
	"html/template"
	"net/http"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/alextotalk/tc-chat/internal/service"
)

type Handler struct {
	services *service.Service
	mux      *http.ServeMux
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
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

func (h *Handler) initRoutes() {
	h.mux.HandleFunc("/", indexHandler)
}

func (h *Handler) NewRouter() *http.ServeMux {
	h.initRoutes()
	return h.mux
}
