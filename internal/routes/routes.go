package routes

import (
	"github.com/gorilla/mux"
	"github.com/gscaramuzzino/go-ecommerce/internal/handlers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	// Product routes
	r.HandleFunc("/products", handlers.ListProducts).Methods("GET")
	r.HandleFunc("/products/{productID}", handlers.GetProductDetails).Methods("GET")

	// Cart routes
	r.HandleFunc("/cart", handlers.GetCart).Methods("GET")
	r.HandleFunc("/cart/{productID}", handlers.AddToCart).Methods("POST")

	// Product routes
	r.HandleFunc("/users/{userID}/checkout", handlers.Checkout).Methods("POST")

	return r
}
