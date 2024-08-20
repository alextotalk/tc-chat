package handler

import (
	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/alextotalk/tc-chat/internal/service"
	"html/template"
	"net/http"
)

type Handler struct {
	services *service.Service
	mux      *http.ServeMux
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
		mux:      http.NewServeMux(),
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

func authHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("templates/auth.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		return
	} // Write to response writer
}

func (h *Handler) initRoutes() {
	h.mux.HandleFunc("/", indexHandler)
	h.mux.HandleFunc("/auth", authHandler)
}

func (h *Handler) NewRouter() *http.ServeMux {
	h.initRoutes()
	return h.mux
}
