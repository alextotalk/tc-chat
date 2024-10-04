package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/alextotalk/tc-chat/internal/service"
	"github.com/golang-jwt/jwt/v5"
)

type ProfileHandler struct {
	*Handler
	service *service.ProfileService
	auth    *service.AuthService
}

func (h *ProfileHandler) editProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := service.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return claims, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Authorized with email %s!", claims.Email)))

	tpl := template.Must(template.ParseFiles("templates/edit-profile.html"))
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		err := h.service.EditProfile(ctx, claims.Email, name, email, phone, password)
		if err != nil {
			log.Println("Something went wrong: ", err)
		}
		tpl.Execute(w, struct{ Success bool }{true})
		// Successful login, redirect to a welcome page.
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

// TODO: write this method with template or dialog window
func (h *ProfileHandler) deleteProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := service.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return claims, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Authorized with email %s!", claims.Email)))

	tpl := template.Must(template.ParseFiles("templates/delete-profile.html"))
	if r.Method == http.MethodDelete {
		err := h.service.DeleteProfile(ctx, claims.Email)
		if err != nil {
			log.Println("Something went wrong: ", err)
		}
		tpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}

}

func (h *ProfileHandler) initRoutes() {
	h.mux.HandleFunc("/edit", h.editProfile)
	h.mux.HandleFunc("/login", h.deleteProfile)
}
