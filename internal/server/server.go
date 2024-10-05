package server

import (
	"context"
	"net/http"
)

type Server struct {
	mux    *http.ServeMux
	server *http.Server
	router  
}

func NewServer(handler *http.ServeMux) *Server {
	return &Server{
		server: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
		router := mux.NewRouter()

 // Define the endpoints for CRUD operations
 router.HandleFunc("/items", getItems).Methods("GET")
 router.HandleFunc("/items/{id}", getItem).Methods("GET")
 router.HandleFunc("/items", createItem).Methods("POST")
 router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
 router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE"
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
