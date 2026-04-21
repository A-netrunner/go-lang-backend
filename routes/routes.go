package routes

import (
	"fmt"
	"net/http"

	"github.com/A-netrunner/go-backend/handlers"
	"github.com/A-netrunner/go-backend/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	//prefix api
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello there")
	})
	//example URL: http://localhost:8080/api/users
	
	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.Use(middleware.Logger)
	return r
}
