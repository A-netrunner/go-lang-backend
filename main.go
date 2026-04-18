package main

import (
	"log"
	"net/http"

	"github.com/A-netrunner/go-backend/routes"
)

func main() {
	r := routes.NewRouter()

	log.Println("server running on http://localhost:8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
