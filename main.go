package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

// EntryPoint Of Our App
func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define a route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Log that the server is running
	log.Println("Server is running on port", PORT)

	// Start the server
	http.ListenAndServe(PORT, r)
}
