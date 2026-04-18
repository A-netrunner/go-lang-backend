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
