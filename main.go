package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const PORT string = ":8080"

// Product Struct
type Product struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Price         string    `json:"price"`
	Category      string    `json:"category"`
	ProductNumber int       `json:"productNumber"`
}

func initRoutes(r *mux.Router) {
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/products", ProductsHandler).Methods(http.MethodGet)
}

func JSONEncoder(w http.ResponseWriter, data interface{}, statusCode int) {
	// Set Content Type and Status Code
	w.Header().Set("Content-Type", "application/json")
	// Set Status Code
	w.WriteHeader(statusCode)
	// Encode Data to JSON and Write to Response
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Create Mux Router and Initialize Routes
	r := mux.NewRouter()

	// Initialize Routes
	initRoutes(r)

	log.Printf("Starting Server on Port %s", PORT)
	// Start Server
	log.Fatal(http.ListenAndServe(PORT, r))
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey, API!\n"))
}

/*
ProductsHandler - Handler for Products
- GET /products - Get All Products
*/
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: uuid.New(), Name: "Product 1", Price: "100", Category: "Category 1"},
		{ID: uuid.New(), Name: "Product 2", Price: "200", Category: "Category 2"},
		{ID: uuid.New(), Name: "Product 3", Price: "300", Category: "Category 3"},
	}

	JSONEncoder(w, products, http.StatusOK)
}
