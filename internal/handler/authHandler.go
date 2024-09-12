package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/alextotalk/tc-chat/internal/service"
)

type AuthHandler struct {
	*Handler
	service *service.AuthService
}

func (h *AuthHandler) singUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tpl := template.Must(template.ParseFiles("templates/auth.html"))
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		phone := r.FormValue("phone")

		err := h.service.Singup(ctx, name, email, password, phone)
		if err != nil {
		}
		tpl.Execute(w, struct{ Success bool }{true})
		// Successful login, redirect to a welcome page.
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/login.html"))
	ctx := r.Context()
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		_, err := h.service.LoginUser(ctx, email, password)
		if err != nil {
			log.Println("Something went wrong: ", err)
		}
		// Successful login, redirect to a welcome page.
		tpl.Execute(w, struct{ Success bool }{true})
		// Successful login, redirect to a welcome page.
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
