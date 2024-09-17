package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// PORT is the port that the server will run on
	PORT = ":8080"
	// VERSION is the version of the API
	VERSION = "v1"
)

func main() {
	// Create Mux Router and Initialize Routes
	r := mux.NewRouter()
	// Initialize Routes
	r.HandleFunc("/api/"+VERSION, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to My API"))
	}).Methods(http.MethodGet)

	log.Printf("Starting Server on Port %s", PORT)
	// Start Server
	log.Fatal(http.ListenAndServe(PORT, r))
}
