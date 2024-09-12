package handler

import (
	"github.com/a-h/templ"
	"github.com/alextotalk/tc-chat/internal/service"
	"github.com/alextotalk/tc-chat/templates/components"
	templates "github.com/alextotalk/tc-chat/templates/home"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(templates.Index()))

	// Група для базових роутів
	//mux.Handle("/", UserMiddleware(http.HandlerFunc(homeHandler)))
	//mux.Handle("/set-name", UserMiddleware(http.HandlerFunc(setNameHandler)))

	// Група для авторизації
	authMux := http.NewServeMux()
	alertHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.Alert("hi").Render(r.Context(), w)
	})
	authMux.HandleFunc("/login", alertHandler)
	//authMux.HandleFunc("/register", registerHandler)
	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))

	// Група для чатів (доступна тільки для авторизованих користувачів)
	//chatMux := http.NewServeMux()
	//chatMux.HandleFunc("/create", createChatHandler)
	//chatMux.HandleFunc("/list", listChatHandler)
	//mux.Handle("/chat/", AuthMiddleware(http.StripPrefix("/chat", chatMux)))
	return mux
}
func loginHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("tokenString"))
}
