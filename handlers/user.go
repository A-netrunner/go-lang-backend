package handlers

import (
	"encoding/json"
	"net/http"

	user "github.com/A-netrunner/go-backend/models"
)

// helper for send json response
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, user.DummyUsers)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u user.Users
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.DummyUsers = append(user.DummyUsers, u)
	writeJSON(w, http.StatusCreated, u)
}
