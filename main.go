package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Write Static HTML File
	http.ServeFile(w, r, "static/index.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Write Static HTML File
	http.ServeFile(w, r, "static/about.html")
}

// EntryPoint Of Our App
func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define a route
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/about", AboutHandler)

	// Log that the server is running
	log.Println("Server is running on port", PORT)

	// Start the server
	http.ListenAndServe(PORT, r)
}
