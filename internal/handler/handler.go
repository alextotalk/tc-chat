package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/alextotalk/tc-chat/internal/service"
	components "github.com/alextotalk/tc-chat/templates/components"
	templates "github.com/alextotalk/tc-chat/templates/home"
)

type Handler struct {
	services *service.Service
	mux      *http.ServeMux
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		mux:      http.NewServeMux(),
		services: services,
	}
}

//
//var tpl = template.Must(template.ParseFiles("templates/index.html"))

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	// Pass any necessary data to the template (optional)
//	users := []domain.User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//	}
//
//	data := map[string]interface{}{
//		"Users": users,
//	}
//	tpl.Execute(w, data) // Write to response writer
//}

func (h *Handler) initRoutes() {
	h.mux.Handle("/", templ.Handler(templates.Index()))
	h.mux.Handle("/login", templ.Handler(components.LogIn()))
	h.mux.Handle("/signup", templ.Handler(components.LogIn()))

}

func (h *Handler) NewRouter() *http.ServeMux {
	h.initRoutes()
	return h.mux
}
